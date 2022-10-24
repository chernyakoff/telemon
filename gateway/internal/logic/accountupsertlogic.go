package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/client"
	"github.com/chernyakoff/telemon/gateway/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AccountUpsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountUpsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountUpsertLogic {
	return &AccountUpsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountUpsertLogic) AccountUpsert(req *client.AccountUpsertRequest) (resp *client.MessageResponse, err error) {
	messageResponse, err := l.svcCtx.AccountRpc.Upsert(l.ctx, req)
	return messageResponse, err
}
