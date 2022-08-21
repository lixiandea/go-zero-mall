package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/noterpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	NoteRpc noterpc.NoteRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		NoteRpc: noterpc.NewNoteRpc(zrpc.MustNewClient(c.NoteRpcConf)),
	}
}
