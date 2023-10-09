package inits

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"log"
)

var Nacos NacosConfig

// var GlobalCon *GlobalConfig
// var Domains = &GlobalCon.DomainConfig
var Domains Domain

func init() {
	viper.SetConfigFile("/Users/jerorry/Documents/gocode/src/2102Amonth/rpc/order_srv/etc/ordersrc.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	Nacos.Server = viper.GetString("Nacos.Server")
	Nacos.Port = viper.GetInt("Nacos.Port")
	Nacos.NameSpaceId = viper.GetString("Nacos.NameSpaceId")
	Nacos.DataId = viper.GetString("Nacos.DataId")
	Nacos.Group = viper.GetString("Nacos.Group")

	log.Println(Nacos)

	GetDomainConfig()
}

func GetDomainConfig() {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(Nacos.NameSpaceId), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	// 创建serverConfig的另一种方式
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(Nacos.Server, uint64(Nacos.Port)),
	}
	// 创建动态配置客户端
	c, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	// 获取配置
	content, _ := c.GetConfig(vo.ConfigParam{
		DataId: Nacos.DataId,
		Group:  Nacos.Group,
	})

	json.Unmarshal([]byte(content), &Domains)

	fmt.Println(Domains.Domain)
}
