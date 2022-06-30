package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/product/rpc/rpc"

	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.ProductRpc.Create(l.ctx, &rpc.CreateRequest{Info: &rpc.ProductInfo{
		Name:   req.Info.Name,
		Desc:   req.Info.Desc,
		Amount: req.Info.Amount,
		Status: req.Info.Status,
		Stock:  req.Info.Stock,
	}})
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
