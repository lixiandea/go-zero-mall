package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/note/note/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

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

func (l *GetNoteLogic) GetNote(in *note.GetNoteRequest) (*note.GetNoteResponse, error) {
	res, err := l.svcCtx.NoteModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "笔记不存在")
		}
		return nil, err
	}
	return &note.GetNoteResponse{
		UserId:  res.Uid,
		Content: res.Content,
	}, nil
}
