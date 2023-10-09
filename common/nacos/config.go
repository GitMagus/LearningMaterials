package main

type Mysql struct {
	MysqlHost string `json:"MysqlHost"`
	MysqlProd int    `json:"MysqlProd"`
	DbName    string `json:"DbName"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
}

type AliPay struct {
	APPID                string `json:"APPID"`
	PrivateKey           string `json:"PrivateKey"`
	PublicKey            string `json:"PublicKey"`
	AlipayGatewayAddress string `json:"AlipayGatewayAddress"`
}

type ElasticSearch struct {
	EsHost string `json:"EsHost"`
	EsProd int    `json:"EsProd"`
}

type Rocketmq struct {
	RocketmqHost string `json:"RocketmqHost"`
	RocketmqProd string `json:"RocketmqProd"`
}

type NaCosConfig struct {
	Mysql         `json:"Mysql"`
	AliPay        `json:"AliPay"`
	Rocketmq      `json:"Rocketmq"`
	ElasticSearch `json:"ElasticSearch"`
}

var NaCos NaCosConfig
