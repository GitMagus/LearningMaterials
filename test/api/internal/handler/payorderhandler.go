package handler

import (
	"net/http"

	"2102Amonth/api/internal/logic"
	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPayOrderLogic(r.Context(), svcCtx)
		resp, err := l.PayOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
