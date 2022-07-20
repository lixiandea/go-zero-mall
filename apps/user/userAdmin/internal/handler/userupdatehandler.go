package handler

import (
	"net/http"

	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/logic"
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/user/userAdmin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserUpdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
