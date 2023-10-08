package svc

import (
	"2102Aweek3/bff/internal/config"
	"2102Aweek3/integral-srv/integralclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	IntegralRpc integralclient.Integral
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		IntegralRpc: integralclient.NewIntegral(zrpc.MustNewClient(c.IntegralRpc)),
	}
}
