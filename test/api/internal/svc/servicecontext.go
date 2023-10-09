package svc

import (
	"2102Amonth/api/internal/config"
	"2102Amonth/rpc/order_srv/ordersrv"
	"github.com/zeromicro/go-zero/zrpc"
	"inventory_srv/inventorysrvclient"
)

type ServiceContext struct {
	Config          config.Config
	OrderSrvRpc     ordersrv.OrderSrv
	InventorySrvRpc inventorysrvclient.InventorySrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		OrderSrvRpc:     ordersrv.NewOrderSrv(zrpc.MustNewClient(c.OrderSrvRpc)),
		InventorySrvRpc: inventorysrvclient.NewInventorySrv(zrpc.MustNewClient(c.InventorySrvRpc)),
	}
}
