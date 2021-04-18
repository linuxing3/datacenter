package handler

import (
	"net/http"

	logic "datacenter/internal/logic/sms/flashpromotion"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func FlashPromotionDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteFlashPromotionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFlashPromotionDeleteLogic(r.Context(), ctx)
		resp, err := l.FlashPromotionDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
