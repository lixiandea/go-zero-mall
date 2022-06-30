package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/order/order/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/order/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *order.RemoveRequest) (*order.RemoveResponse, error) {
	_, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, err
	}
	err = l.svcCtx.OrderModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &order.RemoveResponse{}, nil
}
