package logic

import (
	"context"

	"github.com/chernyakoff/telemon/account/internal/svc"
	"github.com/chernyakoff/telemon/account/types/types"

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

	exists := l.svcCtx.AccountModel.Exists(l.ctx, req.Phone)

	dto := &model.Account{}

	helper.Recast(req, dto)

	if exists {
		l.svcCtx.AccountModel.Update(exists.id)
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", "Эккаунт с таким телефоном уже существует"), "")
	}

	_, err = l.svcCtx.AccountModel.Insert(l.ctx, dto)

	if err != nil {
		return nil, errorx.NewCodeError(202, fmt.Sprintf("%v", err), "")
	}

	return &types.ResponseMessage{
		Success: true,
		Message: "Эккаунт добавлен",
	}, nil
	// todo: add your logic here and delete this line

	return &types.MessageResponse{}, nil
}
