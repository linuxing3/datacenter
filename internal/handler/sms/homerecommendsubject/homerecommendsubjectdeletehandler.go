package handler

import (
	"net/http"

	logic "datacenter/internal/logic/sms/homerecommendsubject"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HomeRecommendSubjectDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteHomeRecommendSubjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHomeRecommendSubjectDeleteLogic(r.Context(), ctx)
		resp, err := l.HomeRecommendSubjectDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
