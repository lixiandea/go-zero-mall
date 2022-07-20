package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"
	cryptx "github.com/lixiandea/go-zero-mall/common/crypt"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user_rpc.UserLoginRequest) (*user_rpc.UserLoginResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "账号不存在")
		}
	}
	if user.Password != cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.UserPassword) {
		return nil, status.Error(403, "账号密码错误")
	}

	return &user_rpc.UserLoginResponse{UserId: user.Id, Info: &user_rpc.UserInfo{
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
