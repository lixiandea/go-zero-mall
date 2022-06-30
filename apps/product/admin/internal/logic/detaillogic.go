package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/product/rpc/rpc"

	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.ProductRpc.Detail(l.ctx, &rpc.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DetailResponse{Info: types.ProductInfo{
		Name:   res.Info.Name,
		Desc:   res.Info.Desc,
		Stock:  res.Info.Stock,
		Amount: res.Info.Amount,
		Status: res.Info.Status,
	}}, nil
}
