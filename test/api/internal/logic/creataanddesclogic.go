package logic

import (
	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"
	"2102Amonth/rpc/order_srv/ordersrv"
	"context"
	"inventory_srv/inventorysrvclient"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
)

var dtmServer = "etcd://localhost:2379/dtmservice"

type CreataAndDescLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreataAndDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreataAndDescLogic {
	return &CreataAndDescLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreataAndDescLogic) CreataAndDesc(req *types.CDRequest) (resp *types.CDResponse, err error) {
	// 将下单动作和减库存以分布式事务的方式来实现
	orderServer, _ := l.svcCtx.Config.OrderSrvRpc.BuildTarget()
	inventoryServe, _ := l.svcCtx.Config.InventorySrvRpc.BuildTarget()

	createRequest := ordersrv.CreateRequest{
		Uid:     int64(req.Uid),
		Gid:     int64(req.Gid),
		Count:   int64(req.Count),
		Fee:     float32(req.Fee),
		Subject: req.Subject,
	}
	inventoryRequest := inventorysrvclient.DRequest{
		Gid:   int64(req.Gid),
		Count: int64(req.Count),
	}
	genid := dtmgrpc.MustGenGid(dtmServer)
	sage := dtmgrpc.NewSagaGrpc(dtmServer, genid).
		Add(orderServer+"/order_src.Order_Srv/CreateOrder", orderServer+"/order_src.Order_Srv/RollbackCreateOrder", &createRequest).
		Add(inventoryServe+"/inventory_srv.Inventory_srv/DescInventory", inventoryServe+"/inventory_srv.Inventory_srv/AddInventory", &inventoryRequest)
	err = sage.Submit()
	if err != nil {
		return nil, err
	}
	return
}
