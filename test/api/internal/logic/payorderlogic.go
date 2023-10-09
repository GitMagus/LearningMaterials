package logic

import (
	"2102Amonth/common/database"
	pay_service "2102Amonth/common/pay-service"
	"context"

	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req *types.PRequest) (resp *types.PResponse, err error) {
	orderNo := req.OrderNo

	orderInfo := database.GetOrderInfo(orderNo)
	// 调用第三方，生成支付url
	payUrl := pay_service.CreatePayUrl(orderInfo)

	return &types.PResponse{
		Code:   200,
		Msg:    "success",
		PayUrl: payUrl,
	}, nil
}
