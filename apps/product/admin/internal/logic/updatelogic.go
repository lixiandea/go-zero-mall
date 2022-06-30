package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/product/rpc/rpc"

	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/types"

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
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &rpc.UpdateRequest{
		Id: req.Id,
		Info: &rpc.ProductInfo{
			Name:   req.Info.Name,
			Desc:   req.Info.Desc,
			Amount: req.Info.Amount,
			Status: req.Info.Status,
			Stock:  req.Info.Stock,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil

}
