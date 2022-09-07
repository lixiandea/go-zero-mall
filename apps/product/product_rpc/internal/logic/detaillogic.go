package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, err
	}

	return &product.DetailResponse{
		Id:     int64(res.Id),
		Stock:  int64(res.Stock),
		Status: int64(res.Status),
		Desc:   res.Desc,
		Amount: int64(res.Amount),
		Name:   res.Name,
	}, nil
}
