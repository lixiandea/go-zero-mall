package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	NoteRpc note.Note
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		NoteRpc: note.NewNote(zrpc.MustNewClient(c.NoteRpcConf)),
	}
}
