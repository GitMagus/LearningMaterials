package handler

import (
	pay_service "2102Amonth/common/pay-service"
	"fmt"
	"net/http"

	"2102Amonth/api/internal/logic"
	"2102Amonth/api/internal/svc"
	"2102Amonth/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncNotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NRequest
		r.ParseForm()
		var noti, err = pay_service.Client.DecodeNotification(r.Form)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSyncNotifyLogic(r.Context(), svcCtx)
		resp, err := l.SyncNotify(&req, noti)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			fmt.Println("success")
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
