// Code generated by goctl. DO NOT EDIT.
// Source: integral.proto

package integralclient

import (
	"context"

	"2102Aweek3/integral-srv/pb/integral"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IRequest  = integral.IRequest
	IResponse = integral.IResponse

	Integral interface {
		// 随机获取积分
		SendIntegral(ctx context.Context, in *IRequest, opts ...grpc.CallOption) (*IResponse, error)
	}

	defaultIntegral struct {
		cli zrpc.Client
	}
)

func NewIntegral(cli zrpc.Client) Integral {
	return &defaultIntegral{
		cli: cli,
	}
}

// 随机获取积分
func (m *defaultIntegral) SendIntegral(ctx context.Context, in *IRequest, opts ...grpc.CallOption) (*IResponse, error) {
	client := integral.NewIntegralClient(m.cli.Conn())
	return client.SendIntegral(ctx, in, opts...)
}
