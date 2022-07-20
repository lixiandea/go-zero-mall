package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *user_rpc.UserUpdateRequest) (*user_rpc.UserUpdateResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Info.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "查无此人")
		}
	}
	user.Phone = in.Info.UserPhone
	user.Address = in.Info.UserAddress
	// user.Password = in.Info.UserPassword
	user.Avatar = in.Info.UserAvatar
	user.Email = in.Info.UserEmail
	user.Status = in.Info.UserStatus
	user.Gender = int64(in.Info.UserGender)
	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &user_rpc.UserUpdateResponse{Info: in.Info}, nil
}
