package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoteLogic {
	return &GetNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNoteLogic) GetNote(in *note_rpc.GetNoteRequest) (*note_rpc.GetNoteResponse, error) {
	noteDetail, err := l.svcCtx.NoteModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &note_rpc.GetNoteResponse{
		Title:   noteDetail.Title,
		Context: noteDetail.Content,
		Author:  noteDetail.Author,
	}, nil
}
