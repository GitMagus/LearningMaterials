package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	_ "github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	_ "github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	_ "github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/zeromicro/go-zero/core/logx"
)

func init() {
	// 创建serverConfig的另一种方式
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"localhost",
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	// 创建clientConfig的另一种方式
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("dev"), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	// 获取配置：GetConfig
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: "month",
		Group:  "dev"})
	fmt.Println(content)
	JsonUnmarshal(content)

	// 监听配置变化：ListenConfig
	err := configClient.ListenConfig(vo.ConfigParam{
		DataId: "dataId",
		Group:  "group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("naCos配置文件发生变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			JsonUnmarshal(data)
		},
	})
	if err != nil {
		logx.Error("监听配置文件变化失败")
		return
	}
	fmt.Println("NaCos序列化成功")
	fmt.Println(NaCos)
}
func JsonUnmarshal(data string) {
	err := json.Unmarshal([]byte(data), &NaCos)
	if err != nil {
		logx.Error("naCos 反序列化 失败")
		return
	}
	logx.Info("naCos 反序列化 ok")
}
