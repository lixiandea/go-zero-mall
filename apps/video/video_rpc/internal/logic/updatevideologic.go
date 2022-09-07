package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/model"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/pb/video_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVideoLogic {
	return &UpdateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVideoLogic) UpdateVideo(in *video_rpc.UpdateVideoRequest) (*video_rpc.UpdateVideoResponse, error) {
	info, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.VideoModel.Update(l.ctx, &model.Video{
		Id:          info.Id,
		Description: info.Description,
		Url:         in.URL,
		Author:      in.Author,
		Title:       in.Title,
	})
	return &video_rpc.UpdateVideoResponse{}, nil
}
