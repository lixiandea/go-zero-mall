package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/noteadmin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NoteadminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewNoteadminLogic(r.Context(), svcCtx)
		resp, err := l.Noteadmin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
