package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/pb/video_rpc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteVideoLogic) DeleteVideo(in *video_rpc.DeleteVideoRequest) (*video_rpc.DeleteVideoResponse, error) {
	video_info, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if video_info == nil {
		return nil, status.Error(404, "video info not found")
	}
	uid := l.ctx.Value("uid").(int)

	if video_info.Author != int64(uid) {
		return nil, status.Error(403, " authorization needed ")
	}
	err = l.svcCtx.VideoModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &video_rpc.DeleteVideoResponse{}, nil
}
