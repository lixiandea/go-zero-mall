package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"
	"github.com/lixiandea/go-zero-mall/common/jwtx"
	"time"

	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	res, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &user_rpc.UserLoginRequest{
		UserName:     req.UserName,
		UserPassword: req.Password,
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.UserId)
	if err != nil {
		return nil, err
	}
	Info := &types.UserInfo{
		UserName: res.Info.UserName,
		Avatar:   res.Info.UserAvatar,
		Address:  res.Info.UserAddress,
		Phone:    res.Info.UserPhone,
		Email:    res.Info.UserEmail,
		Gender:   int(res.Info.UserGender),
		Status:   res.Info.UserStatus,
		Desc:     res.Info.UserDesc,
	}
	return &types.LoginResponse{
		User_id:      res.UserId,
		Info:         *Info,
		AccessExpire: accessExpire + now,
		AccessToken:  token,
	}, nil
}
