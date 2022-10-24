package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/common/helper"
	"github.com/chernyakoff/telemon/account/internal/svc"
	types "github.com/chernyakoff/telemon/account/types/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnyLogic {
	return &GetAnyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnyLogic) GetAny(in *types.AccountGetAnyRequest) (*types.AccountGetResponse, error) {
	account, err := l.svcCtx.AccountModel.FindAny(l.ctx)

	response := types.AccountGetResponse{}
	helper.Recast(account, &response)

	return &response, err
}
