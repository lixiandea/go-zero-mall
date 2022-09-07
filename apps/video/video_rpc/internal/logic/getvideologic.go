package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/pb/video_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoLogic) GetVideo(in *video_rpc.GetVideoRequest) (*video_rpc.GetVideoResponse, error) {
	video_info, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &video_rpc.GetVideoResponse{
		Title:  video_info.Title,
		URL:    video_info.Url,
		Author: video_info.Id,
	}, nil
}
