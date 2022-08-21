package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoteLogic {
	return &UpdateNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNoteLogic) UpdateNote(in *note_rpc.UpdateNoteRequest) (*note_rpc.UpdateNoteResponse, error) {
	noteDetail, err := l.svcCtx.NoteModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if noteDetail == nil {
		return nil, status.Error(400, "no article found")
	}
	if noteDetail.Author != l.ctx.Value("uid") {
		return nil, status.Error(403, "not the author which the article belongs to")
	}
	noteDetail.Title = in.Title
	noteDetail.Content = in.Context
	err = l.svcCtx.NoteModel.Update(l.ctx, noteDetail)
	if err != nil {
		return nil, err
	}
	return &note_rpc.UpdateNoteResponse{Id: noteDetail.Id}, nil
}
