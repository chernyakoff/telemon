package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/common/helper"
	"github.com/chernyakoff/telemon/account/internal/svc"
	"github.com/chernyakoff/telemon/account/model"
	types "github.com/chernyakoff/telemon/account/types/account"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertLogic {
	return &UpsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpsertLogic) Upsert(in *types.AccountUpsertRequest) (*types.MessageResponse, error) {
	dto := &model.Account{}
	helper.Recast(in, dto)
	message, err := l.svcCtx.AccountModel.Upsert(l.ctx, dto)
	return &types.MessageResponse{Message: message}, err
}
