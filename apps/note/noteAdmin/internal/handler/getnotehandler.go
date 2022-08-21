package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteAdmin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetNoteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetNoteLogic(r.Context(), svcCtx)
		resp, err := l.GetNote(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
