package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/model"
	"github.com/lixiandea/go-zero-mall/apps/product/product_rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  uint64(in.Stock),
		Amount: uint64(in.Amount),
		Status: uint64(in.Status),
	})
	if err != nil {
		return nil, err
	}
	resID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &product.CreateResponse{Id: resID}, nil
}
