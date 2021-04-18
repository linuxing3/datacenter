package handler

import (
	"net/http"

	logic "datacenter/internal/logic/sms/coupon"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CouponAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCouponAddLogic(r.Context(), ctx)
		resp, err := l.CouponAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
