package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/video/video_admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Video_adminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideo_adminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Video_adminLogic {
	return &Video_adminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Video_adminLogic) Video_admin(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
