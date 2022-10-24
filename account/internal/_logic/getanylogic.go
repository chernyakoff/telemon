package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/internal/svc"
	"github.com/chernyakoff/telemon/account/types/types"

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
	// todo: add your logic here and delete this line

	return &types.AccountGetResponse{}, nil
}
