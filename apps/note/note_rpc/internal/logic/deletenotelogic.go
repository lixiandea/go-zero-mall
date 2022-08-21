package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DeleteNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoteLogic {
	return &DeleteNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteNoteLogic) DeleteNote(in *note_rpc.DeleteNoteRequest) (*note_rpc.DeleteNoteResponse, error) {
	noteDetail, err := l.svcCtx.NoteModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if noteDetail.Author != l.ctx.Value("uid") {
		return nil, status.Error(403, "not the author which the article belongs to")
	}
	err = l.svcCtx.NoteModel.Delete(l.ctx, noteDetail.Id)
	if err != nil {
		return nil, err
	}
	return &note_rpc.DeleteNoteResponse{}, nil
}
