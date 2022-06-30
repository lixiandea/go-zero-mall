package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoteLogic {
	return &UpdateNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNoteLogic) UpdateNote(req *types.UpdateNoteRequest) (resp *types.UpdateNoteResponse, err error) {
	_, err = l.svcCtx.NoteRpc.UpdateNote(l.ctx, &note.UpdateNoteRequest{
		Id:      req.Id,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateNoteResponse{}, nil
}
