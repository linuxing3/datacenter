package handler

import (
	"net/http"

	logic "datacenter/internal/logic/member/member"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MemberAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberAddLogic(r.Context(), ctx)
		resp, err := l.MemberAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
