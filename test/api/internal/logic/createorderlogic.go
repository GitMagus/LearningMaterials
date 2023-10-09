package logic

import (
	"2102Amonth/rpc/order_srv/pb/order_src"
	"context"
	"log"

	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	create, _ := l.svcCtx.OrderSrvRpc.CreateOrder(l.ctx, &order_src.CreateRequest{
		Uid:     int64(req.Uid),
		Gid:     int64(req.Gid),
		Fee:     float32(req.Fee),
		Count:   int64(req.Count),
		Subject: req.Subject,
	})

	log.Println(create)
	return &types.Response{
		Code: 200,
		Msg:  "success",
		Data: types.OrderInfo{
			OrderNo: create.OrderNo,
		},
	}, nil
}
