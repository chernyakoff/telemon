package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/common/helper"
	"github.com/chernyakoff/telemon/account/internal/svc"
	types "github.com/chernyakoff/telemon/account/types/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *types.AccountGetRequest) (*types.AccountGetResponse, error) {

	account, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.Phone)

	response := types.AccountGetResponse{}
	helper.Recast(account, &response)

	return &response, err
}
