package logic

import (
	"2102Aweek3/integral-srv/integralclient"
	"context"
	"strconv"

	"2102Aweek3/bff/internal/svc"
	"2102Aweek3/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserIntegralLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserIntegralLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserIntegralLogic {
	return &UpdateUserIntegralLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserIntegralLogic) UpdateUserIntegral(req *types.Request) (resp *types.Response, err error) {
	userId := req.UserId

	res, _ := l.svcCtx.IntegralRpc.SendIntegral(l.ctx, &integralclient.IRequest{
		UserId: strconv.Itoa(userId),
	})
	if res.Result == false {
		return &types.Response{
			Code: -1,
			Msg:  "",
		}, nil
	}
	return &types.Response{
		Code: 200,
		Msg:  "success",
	}, nil
}
