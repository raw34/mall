package logic

import (
	"context"
	"errors"

	"raw34/mall/order/api/internal/svc"
	"raw34/mall/order/api/internal/types"
	"raw34/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}

	if user.Name != "test" {
		return nil, errors.New("用户不存在")
	}

	return &types.OrderReply{
		Id:   req.Id,
		Name: "test order",
	}, nil
}
