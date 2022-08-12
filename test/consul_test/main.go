package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	//check := &api.AgentServiceCheck{
	//	HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
	//	Timeout:                        "5s",
	//	Interval:                       "5s",
	//	DeregisterCriticalServiceAfter: "10s",
	//}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	//registration.Check = check

	if err := client.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}
	return nil
}

func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	registerClient := NewRegistryClient("82.157.166.247", 8500)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register("82.157.166.247", 8022, "user-web1", []string{"user","wei"}, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//设置要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		if err:= registerClient.DeRegister(serviceId);err != nil{
			zap.S().Info("注销失败：", err.Error())
		}else {
			done <- true
			fmt.Println("注销成功")
			zap.S().Info("注销成功")
		}
	}()

	fmt.Println("等待信号")
	<-done
	fmt.Println("进程被终止")

}
