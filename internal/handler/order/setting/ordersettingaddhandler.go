package handler

import (
	logic "datacenter/internal/logic/order/setting"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func OrderSettingAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOrderSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOrderSettingAddLogic(r.Context(), ctx)
		resp, err := l.OrderSettingAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
