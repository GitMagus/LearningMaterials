package logic

import (
	"2102Aweek3/integral-srv/internal/svc"
	"2102Aweek3/integral-srv/pb/integral"
	"2102Aweek3/rocketmq/publish"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendIntegralLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendIntegralLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendIntegralLogic {
	return &SendIntegralLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 随机获取积分
func (l *SendIntegralLogic) SendIntegral(in *integral.IRequest) (*integral.IResponse, error) {
	// 接收参数
	userId := in.UserId
	newIntegral := in.Integral
	optrationType := in.Type
	body := map[string]interface{}{
		"user_id":  userId,
		"integral": newIntegral,
		"type":     optrationType,
	}
	b, _ := json.Marshal(body)
	// 调用消息生成方法
	r := publish.Publish("score", string(b))
	if r == false {
		return &integral.IResponse{}, nil
	}
	return &integral.IResponse{
		Result: true,
	}, nil
}
