package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/order/model"
	"github.com/lixiandea/go-zero-mall/apps/order/order/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/order/order/order"
	"github.com/lixiandea/go-zero-mall/apps/product/product/product"
	"github.com/lixiandea/go-zero-mall/apps/user/user/user"
	"google.golang.org/grpc/status"

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

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {

	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{Id: in.Pid})
	if err != nil {
		return nil, err
	}
	if res.Info.Stock <= 0 {
		return nil, status.Error(100, "库存不足")
	}
	newOrder := &model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: res.Info.Amount,
		Status: res.Info.Status,
	}
	insertRes, err := l.svcCtx.OrderModel.Insert(l.ctx, newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newOrder.Id, err = insertRes.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{Id: newOrder.Id}, nil
}
