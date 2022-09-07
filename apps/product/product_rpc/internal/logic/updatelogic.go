package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/types/product"

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

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}
	res.Amount = uint64(in.Amount)
	res.Status = uint64(in.Status)
	res.Stock = uint64(in.Stock)
	res.Name = in.Name
	res.Desc = in.Desc
	err = l.svcCtx.ProductModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &product.UpdateResponse{}, nil
}
