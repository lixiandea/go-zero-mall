package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/product/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/product/rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/rpc/rpc"

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

func (l *UpdateLogic) Update(in *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Info.Name != "" {
		res.Name = in.Info.Name
	}
	if in.Info.Desc != "" {
		res.Desc = in.Info.Desc
	}
	if in.Info.Stock != 0 {
		res.Stock = in.Info.Stock
	}
	if in.Info.Amount != 0 {
		res.Amount = in.Info.Amount
	}
	if in.Info.Status != 0 {
		res.Status = in.Info.Status
	}

	err = l.svcCtx.ProductModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &rpc.UpdateResponse{}, nil

}
