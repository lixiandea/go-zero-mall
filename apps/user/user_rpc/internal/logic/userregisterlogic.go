package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/model"
	"github.com/lixiandea/go-zero-mall/apps/user/user_rpc/user_rpc"
	cryptx "github.com/lixiandea/go-zero-mall/common/crypt"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *user_rpc.UserRegisterRequest) (*user_rpc.UserRegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.Info.UserName)
	if err == nil {
		return nil, status.Error(100, "账号已注册")
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	_, err = l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Info.UserPhone)

	if err == nil {
		return nil, status.Error(100, "手机号已注册")
	}
	if err != model.ErrNotFound {
		return nil, err
	}

	user := &model.User{
		Name:     in.Info.UserName,
		Email:    in.Info.UserEmail,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		Phone:    in.Info.UserPhone,
		Address:  in.Info.UserAddress,
		Gender:   int64(in.Info.UserGender),
		Status:   in.Info.UserStatus,
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &user_rpc.UserRegisterResponse{UserId: id}, nil
}
