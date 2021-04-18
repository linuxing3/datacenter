package handler

import (
	"net/http"

	logic "datacenter/internal/logic/member/address"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MemberAddressUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMemberAddressReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberAddressUpdateLogic(r.Context(), ctx)
		resp, err := l.MemberAddressUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
