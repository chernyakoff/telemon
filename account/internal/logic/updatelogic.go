package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/common/helper"
	"github.com/chernyakoff/telemon/account/internal/svc"
	"github.com/chernyakoff/telemon/account/model"
	types "github.com/chernyakoff/telemon/account/types/account"

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
	dto := &model.Account{}
	helper.Recast(in, dto)
	message := ""
	err := l.svcCtx.AccountModel.Update(l.ctx, dto)
	if err == nil {
		message = "Эккаунт успешно обновлен"
	}
	return &types.MessageResponse{Message: message}, err
}
