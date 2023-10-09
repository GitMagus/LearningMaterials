package logic

import (
	"2102Amonth/common/database"
	"context"
	"github.com/smartwalle/alipay/v3"
	"inventory_srv/inventorysrvclient"
	"time"

	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncNotifyLogic {
	return &SyncNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncNotifyLogic) SyncNotify(req *types.NRequest, noti *alipay.Notification) (resp *types.NResponse, err error) {
	// 通过验签操作，对支付订单进行数据修改
	orderNo := noti.OutTradeNo
	// 根据orderNo查询订单数据
	orderInfo := database.GetOrderInfo(orderNo)
	//if orderInfo.Id > 0 {
	//	logx.Debug(noti.TradeStatus)
	//	if noti.TradeStatus == alipay.TradeStatusSuccess {
	tradeNo := noti.TradeNo
	orderInfo.TradeNo = tradeNo
	orderInfo.Status = 1
	orderInfo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	database.UpdateOrder(orderInfo)
	// 扣减库存
	l.svcCtx.InventorySrvRpc.DescInventory(l.ctx, &inventorysrvclient.DRequest{
		Gid:   int64(orderInfo.Gid),
		Count: int64(orderInfo.Count),
	})
	//	} else {
	//		orderInfo.Status = 2
	//		orderInfo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	//		database.UpdateOrder(orderInfo)
	//	}
	//}

	return
}
