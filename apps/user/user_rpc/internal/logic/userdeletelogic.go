package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *user_rpc.UserDeleteRequest) (*user_rpc.UserDeleteResponse, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "查无此人")
		}
		return nil, err
	}
	if err := l.svcCtx.UserModel.Delete(l.ctx, user.Id); err != nil {
		return nil, status.Error(500, "删除失败")
	}
	info := &user_rpc.UserInfo{
		UserName:    user.Name,
		UserAvatar:  user.Avatar,
		UserEmail:   user.Email,
		UserPhone:   user.Phone,
		UserAddress: user.Address,
		UserDesc:    "",
		UserStatus:  user.Status,
		UserGender:  uint32(user.Gender),
	}
	return &user_rpc.UserDeleteResponse{Info: info}, nil
}
