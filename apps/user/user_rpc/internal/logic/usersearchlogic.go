package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"

	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSearchLogic {
	return &UserSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSearchLogic) UserSearch(in *user_rpc.UserSearchRequest) (*user_rpc.UserSearchResponse, error) {
	// TODO: 需要处理分页查询的内容
	users, err := l.svcCtx.UserModel.FindUsersByName(l.ctx, in.UserName)
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
	return &user_rpc.UserSearchResponse{
		Infoes: userInfos,
	}, nil
}
