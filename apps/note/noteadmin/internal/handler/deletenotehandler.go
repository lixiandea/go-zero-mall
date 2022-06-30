package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteNoteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteNoteLogic(r.Context(), svcCtx)
		resp, err := l.DeleteNote(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
