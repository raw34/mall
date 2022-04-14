package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"raw34/mall/order/api/internal/config"
	"raw34/mall/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
