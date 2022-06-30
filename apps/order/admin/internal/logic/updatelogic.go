package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/order/order/order"

	"github.com/lixiandea/go-zero-mall/apps/order/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/order/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &order.UpdateRequest{
		Id:     req.Id,
		Amount: req.Amount,
		Status: req.Status,
		Uid:    req.Uid,
		Pid:    req.Pid,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil

	return
}
