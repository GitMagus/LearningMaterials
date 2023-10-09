package inits

type GlobalConfig struct {
	DomainConfig Domain
}

type NacosConfig struct {
	Server      string
	Port        int
	NameSpaceId string
	DataId      string
	Group       string
}

type Domain struct {
	Domain string `mapstructure:"domain" json:"domain"`
}
