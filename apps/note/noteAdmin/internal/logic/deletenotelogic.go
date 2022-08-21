package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/types"
	"github.com/lixiandea/go-zero-mall/apps/note/note_rpc/note_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoteLogic {
	return &DeleteNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNoteLogic) DeleteNote(req *types.DeleteNoteRequest) (resp *types.DeleteNoteResponse, err error) {
	_, err = l.svcCtx.NoteRpc.DeleteNote(l.ctx, &note_rpc.DeleteNoteRequest{

		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}
	return &types.DeleteNoteResponse{}, nil
}
