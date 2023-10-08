package handler

import (
	"net/http"

	"2102Aweek3/bff/internal/logic"
	"2102Aweek3/bff/internal/svc"
	"2102Aweek3/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserIntegralHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateUserIntegralLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserIntegral(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
