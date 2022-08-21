package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/types"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"

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
	_, err = l.svcCtx.NoteRpc.UpdateNote(l.ctx, &note_rpc.UpdateNoteRequest{
		Id:      req.Id,
		Title:   req.Title,
		Context: req.Context,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateNoteResponse{}, nil
}
