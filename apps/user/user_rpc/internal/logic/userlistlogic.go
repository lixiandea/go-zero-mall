package logic

import (
	"context"

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

	return &user_rpc.UserListResponse{}, nil
}
