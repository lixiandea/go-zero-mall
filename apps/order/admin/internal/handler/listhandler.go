package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/order/admin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/order/admin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/order/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
