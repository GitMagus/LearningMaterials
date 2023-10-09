package logic

import (
	"context"

	"inventory_srv/internal/svc"
	"inventory_srv/pb/inventory_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInventoryLogic {
	return &AddInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加库存
func (l *AddInventoryLogic) AddInventory(in *inventory_srv.AResquest) (*inventory_srv.AResponse, error) {
	// todo: add your logic here and delete this line

	return &inventory_srv.AResponse{}, nil
}
