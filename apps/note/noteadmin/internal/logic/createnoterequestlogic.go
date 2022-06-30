package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNoteRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNoteRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNoteRequestLogic {
	return &CreateNoteRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNoteRequestLogic) CreateNoteRequest(req *types.CreateNoteRequest) (resp *types.CreateNoteResponse, err error) {
	res, err := l.svcCtx.NoteRpc.CreateNote(l.ctx, &note.CreateNoteRequest{
		Author:  req.Author,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateNoteResponse{Id: res.Id}, nil
}
