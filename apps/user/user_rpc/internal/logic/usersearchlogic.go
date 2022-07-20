package logic

import (
	"context"

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
	//user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)
	//if err != nil {
	//	return nil, err
	//}
	return &user_rpc.UserSearchResponse{}, nil
}
