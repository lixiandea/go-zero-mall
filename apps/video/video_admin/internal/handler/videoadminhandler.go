package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/video/video_admin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/video/video_admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/video/video_admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Video_adminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewVideo_adminLogic(r.Context(), svcCtx)
		resp, err := l.Video_admin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
