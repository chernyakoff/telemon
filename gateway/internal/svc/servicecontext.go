package svc

import (
	"github.com/chernyakoff/telemon/account/client"
	"github.com/chernyakoff/telemon/gateway/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc client.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: client.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
	}
}
