package logic

import (
	"context"
	"github.com/lixiandea/go-zero-mall/apps/note/note/model"
	"github.com/lixiandea/go-zero-mall/apps/user/user/user"
	"google.golang.org/grpc/status"

	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNoteLogic {
	return &CreateNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateNoteLogic) CreateNote(in *note.CreateNoteRequest) (*note.CreateNoteResponse, error) {
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Author})
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, err
	}
	res, err := l.svcCtx.NoteModel.Insert(l.ctx, &model.Note{
		Uid:     in.Author,
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &note.CreateNoteResponse{Id: id}, nil
}
