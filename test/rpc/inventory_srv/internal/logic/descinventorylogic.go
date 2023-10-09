package logic

import (
	"2102Amonth/common/database"
	"context"
	"github.com/dtm-labs/client/dtmcli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"inventory_srv/internal/svc"
	"inventory_srv/pb/inventory_srv"

	"github.com/zeromicro/go-zero/core/logx"
)

type DescInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDescInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DescInventoryLogic {
	return &DescInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 扣减库存
func (l *DescInventoryLogic) DescInventory(in *inventory_srv.DRequest) (*inventory_srv.DResponse, error) {
	gid := in.Gid
	count := in.Count

	res := database.DescGoodCount(int(gid), int(count))
	return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	if res == false {
		//return &inventory_srv.DResponse{
		//	Result: false,
		//}, nil

		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	return &inventory_srv.DResponse{
		Result: true,
	}, nil
}
