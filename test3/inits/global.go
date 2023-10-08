package inits

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"

	_ "github.com/spf13/viper"
)

var (
	Nacos NacosConfig
	//Mysql *MysqlCon = MysqlConfig
	MConfig MysqlCon

	RocketServer = "127.0.0.1:9876"
)

func init() {
	viper.SetConfigFile("/Users/jerorry/Documents/gocode/src/2102Aweek3/config/config-dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic("--------------------读取配置文件出错--------------------")
	}
	Nacos.NacosServer = viper.GetString("NaocsConfig.Server")
	Nacos.NacosPort = viper.GetInt("NaocsConfig.Port")
	Nacos.NacosNamespaceid = viper.GetString("NaocsConfig.NamespaceId")
	Nacos.NacosDataId = viper.GetString("NaocsConfig.DataId")
	Nacos.NacosGroup = viper.GetString("NaocsConfig.Group")

	fmt.Println("--------------------获取naocs成功--------------------")
	fmt.Println(Nacos)
	GetConfig()
}

// 读取nacos配置，作为全局配置
func GetConfig() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(Nacos.NacosServer, uint64(Nacos.NacosPort)),
	}

	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(Nacos.NacosNamespaceid), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	// 建立nacos的客户端
	r, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	res, _ := r.GetConfig(
		vo.ConfigParam{
			DataId: Nacos.NacosDataId,
			Group:  Nacos.NacosGroup,
		},
	)
	fmt.Println(res)
	yaml.Unmarshal([]byte(res), &MConfig)
	fmt.Println(MConfig.Mysql.Host)
}
