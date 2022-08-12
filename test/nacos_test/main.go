package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("82.157.166.247", 8848, constant.WithContextPath("/nacos"), constant.WithScheme("http")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("620ebc22-16f5-431b-a6a2-da9079e141fb"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("tmp/nacos/log"),
		constant.WithCacheDir("tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		fmt.Println("--------------------------")
		panic(err)
	}

	//publish config
	//config key=dataId+group+namespaceId
	//_, err = client.PublishConfig(vo.ConfigParam{
	//	DataId:  "test-data",
	//	Group:   "test-group",
	//	Content: "hello world!",
	//})
	//_, err = client.PublishConfig(vo.ConfigParam{
	//	DataId:  "test-data-2",
	//	Group:   "test-group",
	//	Content: "hello world!",
	//})
	//if err != nil {
	//	fmt.Printf("PublishConfig err:%+v \n", err)
	//}
	//time.Sleep(1 * time.Second)

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "user-web",
		Group:  "dev",
	})
	fmt.Println("GetConfig,config :" + content)

	//Listen config change,key=dataId+group+namespaceId.
	//err = client.ListenConfig(vo.ConfigParam{
	//	DataId: "test-data",
	//	Group:  "test-group",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
	//	},
	//})
	//
	//err = client.ListenConfig(vo.ConfigParam{
	//	DataId: "test-data-2",
	//	Group:  "test-group",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
	//	},
	//})
	//
	//time.Sleep(1 * time.Second)
	//
	//_, err = client.PublishConfig(vo.ConfigParam{
	//	DataId:  "test-data",
	//	Group:   "test-group",
	//	Content: "test-listen",
	//})
	//
	//time.Sleep(1 * time.Second)
	//
	//_, err = client.PublishConfig(vo.ConfigParam{
	//	DataId:  "test-data-2",
	//	Group:   "test-group",
	//	Content: "test-listen",
	//})
	//
	//time.Sleep(2 * time.Second)

	//cancel config change
	//err = client.CancelListenConfig(vo.ConfigParam{
	//	DataId: "test-data",
	//	Group:  "test-group",
	//})
	//
	//time.Sleep(1 * time.Second)
	//_, err = client.DeleteConfig(vo.ConfigParam{
	//	DataId: "test-data",
	//	Group:  "test-group",
	//})
	//time.Sleep(1 * time.Second)
	//
	//searchPage, _ := client.SearchConfig(vo.SearchConfigParm{
	//	Search:   "blur",
	//	DataId:   "",
	//	Group:    "",
	//	PageNo:   1,
	//	PageSize: 10,
	//})
	//fmt.Printf("Search config:%+v \n", searchPage)
}
