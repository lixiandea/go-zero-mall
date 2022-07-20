package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"

	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *user_rpc.UserListRequest) (*user_rpc.UserListResponse, error) {
	// todo: add your logic here and delete this line
	users, err := l.svcCtx.UserModel.ListUserByUid(l.ctx, in.CurUserId, 10)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, model.ErrNotFound
	}
	userInfos := make([]*user_rpc.UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		userInfos[i] = &user_rpc.UserInfo{
			UserName:    users[i].Name,
			UserAvatar:  users[i].Avatar,
			UserEmail:   users[i].Email,
			UserPhone:   users[i].Phone,
			UserAddress: users[i].Address,
			UserDesc:    "",
			UserStatus:  users[i].Status,
			UserGender:  uint32(users[i].Gender),
		}
	}
	return &user_rpc.UserListResponse{
		Infoes: userInfos,
	}, nil
}
