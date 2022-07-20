package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	info := &types.UserInfo{
		UserName: res.Info.UserName,
		Email:    res.Info.UserEmail,
		Phone:    res.Info.UserPhone,
		Address:  res.Info.UserAddress,
		Gender:   int(res.Info.UserGender),
		Status:   res.Info.UserStatus,
		Desc:     res.Info.UserDesc,
		Avatar:   res.Info.UserAvatar,
	}
	//copier.Copy(info, res.Info)
	return &types.UserInfoResponse{
		Info: *info,
	}, nil
}
