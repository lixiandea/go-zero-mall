package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/types"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoteHandlersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddNoteHandlersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoteHandlersLogic {
	return &AddNoteHandlersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddNoteHandlersLogic) AddNoteHandlers(req *types.AddNoteRequest) (resp *types.AddNoteResponse, err error) {
	res, err := l.svcCtx.NoteRpc.AddNote(l.ctx, &note_rpc.AddNoteRequest{
		Title:   req.Title,
		Context: req.Context,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddNoteResponse{
		Id: res.Id,
	}, nil

}
