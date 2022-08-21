package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/model"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoteLogic {
	return &AddNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddNoteLogic) AddNote(in *note_rpc.AddNoteRequest) (*note_rpc.AddNoteResponse, error) {
	res, err := l.svcCtx.NoteModel.Insert(l.ctx, &model.Note{
		Title:   in.Title,
		Content: in.Context,
		Author:  l.ctx.Value("uid").(int64),
	})
	if err != nil {
		return nil, err
	}
	resId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &note_rpc.AddNoteResponse{
		Id: resId,
	}, nil
}
