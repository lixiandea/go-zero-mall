package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddNoteHandlersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddNoteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddNoteHandlersLogic(r.Context(), svcCtx)
		resp, err := l.AddNoteHandlers(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
