package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoteLogic {
	return &GetNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoteLogic) GetNote(req *types.GetNoteRequest) (resp *types.GetNoteResponse, err error) {
	res, err := l.svcCtx.NoteRpc.GetNote(l.ctx, &note.GetNoteRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.GetNoteResponse{
		Author:  res.UserId,
		Content: res.Content,
		Id:      0,
	}, nil
}
