package logic

import (
	"context"

	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_rpc/model"
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
	uid := l.ctx.Value("uid").(int)
	// if err != nil {
	// 	return nil, err
	// }
	result, err := l.svcCtx.VideoModel.Insert(l.ctx, &model.Video{
		Title:       in.Title,
		Description: in.Context,
		Author:      int64(uid),
		Url:         "randomurl",
	})
	if err != nil {
		return nil, err
	}
	vid, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &video_rpc.AddVideoResponse{
		Id: vid,
	}, nil
}
