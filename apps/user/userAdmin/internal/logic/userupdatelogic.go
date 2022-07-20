package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateRequest) (resp *types.UserUpdateResponse, err error) {
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
	_, err = l.svcCtx.UserRpc.UserUpdate(l.ctx, &user_rpc.UserUpdateRequest{
		UserId: req.UserId,
		Info:   info,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserUpdateResponse{
		UserId: req.UserId,
		Info:   req.Info,
	}, nil
}
