package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/pb/video_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoLogic {
	return &AddVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVideoLogic) AddVideo(in *video_rpc.AddVideoRequest) (*video_rpc.AddVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &video_rpc.AddVideoResponse{}, nil
}
