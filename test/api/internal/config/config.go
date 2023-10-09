package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OrderSrvRpc     zrpc.RpcClientConf
	InventorySrvRpc zrpc.RpcClientConf
}
