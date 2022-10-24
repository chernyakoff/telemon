package logic

import (
	"context"

	"github.com/chernyakoff/telemon/gateway/internal/svc"
	"github.com/chernyakoff/telemon/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.MessageResponse, err error) {
	return &types.MessageResponse{
		Message: "pong",
	}, nil
}
