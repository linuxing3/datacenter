package handler

import (
	logic "datacenter/internal/logic/order/setting"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func OrderSettingUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateOrderSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOrderSettingUpdateLogic(r.Context(), ctx)
		resp, err := l.OrderSettingUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
