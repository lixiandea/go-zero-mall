package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	info := &user_rpc.UserInfo{
		UserName:    req.Info.UserName,
		UserAvatar:  req.Info.Avatar,
		UserEmail:   req.Info.Email,
		UserPhone:   req.Info.Phone,
		UserAddress: req.Info.Address,
		UserDesc:    req.Info.Desc,
		UserStatus:  req.Info.Status,
		UserGender:  uint32(req.Info.Gender),
	}
	logx.Info(l.ctx, "UserRegister", "info", info)
	res, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &user_rpc.UserRegisterRequest{Info: info, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		UserId: res.UserId,
	}, nil
}
