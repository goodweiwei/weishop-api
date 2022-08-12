package main


import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//设置要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("等待信号")
	<-done
	fmt.Println("进程被终止")
}