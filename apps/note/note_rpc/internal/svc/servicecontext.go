package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	NoteModel model.NoteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MysqlDataSource.DataSource)
	return &ServiceContext{
		Config:    c,
		NoteModel: model.NewNoteModel(conn, c.RedisConf),
	}
}
