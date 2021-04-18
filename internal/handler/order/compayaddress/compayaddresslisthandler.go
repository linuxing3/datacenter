package handler

import (
	logic "datacenter/internal/logic/order/compayaddress"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CompayAddressListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCompayAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCompayAddressListLogic(r.Context(), ctx)
		resp, err := l.CompayAddressList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
