package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/user/admin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/user/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
