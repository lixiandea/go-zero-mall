package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/types/product"

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

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	err := l.svcCtx.ProductModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}
	return &product.RemoveResponse{}, nil
}
