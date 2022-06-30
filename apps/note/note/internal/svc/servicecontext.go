package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/note/note/model"
	"github.com/lixiandea/go-zero-mall/apps/user/user/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	NoteModel model.NoteModel
	UserRpc   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DateSource)
	return &ServiceContext{
		Config:    c,
		NoteModel: model.NewNoteModel(conn, c.RedisCacheConf),
		UserRpc:   user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
