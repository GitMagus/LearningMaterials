package inits

type NacosConfig struct {
	NacosServer      string
	NacosPort        int
	NacosNamespaceid string
	NacosDataId      string
	NacosGroup       string
}
type MysqlCon struct {
	Mysql MysqlConfig `mapstructure:"Mysql" yaml:"Mysql"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"Host" yaml:"Host"`
	Port     int    `mapstructure:"Port" yaml:"Port"`
	Username string `mapstructure:"Username" yaml:"Username"`
	Password string `mapstructure:"Password" yaml:"Password"`
	Database string `mapstructure:"Database" yaml:"Database"`
}
