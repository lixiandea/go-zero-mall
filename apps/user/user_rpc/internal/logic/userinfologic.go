package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user_rpc.UserInfoRequest) (*user_rpc.UserInfoResponse, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "账号不存在")
		}
	}

	return &user_rpc.UserInfoResponse{Info: &user_rpc.UserInfo{
		UserName:    user.Name,
		UserAvatar:  user.Avatar,
		UserEmail:   user.Email,
		UserPhone:   user.Phone,
		UserAddress: user.Address,
		UserDesc:    "",
		UserStatus:  user.Status,
		UserGender:  uint32(user.Gender),
	}}, nil
}
