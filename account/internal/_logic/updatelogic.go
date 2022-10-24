package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/internal/svc"
	"github.com/chernyakoff/telemon/account/types/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *types.AccountUpdateRequest) (*types.MessageResponse, error) {
	// todo: add your logic here and delete this line

	return &types.MessageResponse{}, nil
}
