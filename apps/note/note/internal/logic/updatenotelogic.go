package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/user/model"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

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

func (l *UpdateNoteLogic) UpdateNote(in *note.UpdateNoteRequest) (*note.UpdateNoteResponse, error) {
	res, err := l.svcCtx.NoteModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "没有找到这个笔记")
		}
	}
	res.Content = in.Content
	res.Uid = in.Id
	err = l.svcCtx.NoteModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &note.UpdateNoteResponse{Id: res.Id}, nil
}
