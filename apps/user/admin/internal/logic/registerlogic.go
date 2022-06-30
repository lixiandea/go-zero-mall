package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/admin/internal/types"
	"github.com/lixiandea/go-zero-mall/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		UserInfo: types.UserInfoResponse{
			ID:     res.Id,
			Name:   res.Name,
			Gender: res.Gender,
			Mobile: res.Mobile,
		},
	}, nil
}
