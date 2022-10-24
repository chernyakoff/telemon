package svc

import (
	"github.com/chernyakoff/telemon/account/internal/config"
	"github.com/chernyakoff/telemon/account/model"

	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel model.AccountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.Pg.Datasource)
	return &ServiceContext{
		Config:       c,
		AccountModel: model.NewAccountModel(conn),
	}
}
