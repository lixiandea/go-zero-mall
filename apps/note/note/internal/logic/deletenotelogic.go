package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/note/note/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *DeleteNoteLogic) DeleteNote(in *note.DeleteNoteRequest) (*note.DeleteNoteResponse, error) {
	res, err := l.svcCtx.NoteModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "笔记不存在")
		}
		return nil, err
	}
	err = l.svcCtx.NoteModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, err
	}
	return &note.DeleteNoteResponse{}, nil
}
