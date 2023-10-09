package logic

import (
	"context"

	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AsyncNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsyncNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AsyncNotifyLogic {
	return &AsyncNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsyncNotifyLogic) AsyncNotify(req *types.NRequest) (resp *types.NResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
