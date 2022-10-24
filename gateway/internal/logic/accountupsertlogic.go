package logic

import (
	"context"

	"github.com/chernyakoff/telemon/gateway/internal/svc"
	"github.com/chernyakoff/telemon/gateway/internal/types"

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

func (l *AccountUpsertLogic) AccountUpsert(req *types.AccountUpsertRequest) (resp *types.MessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
