package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/product/rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/product/rpc/rpc"

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

func (l *RemoveLogic) Remove(in *rpc.RemoveRequest) (*rpc.RemoveResponse, error) {
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	err = l.svcCtx.ProductModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &rpc.RemoveResponse{}, nil
}
