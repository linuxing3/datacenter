package handler

import (
	"net/http"

	logic "datacenter/internal/logic/member/integrationconsumesetting"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func IntegrationConsumeSettingDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteIntegrationConsumeSettingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewIntegrationConsumeSettingDeleteLogic(r.Context(), ctx)
		resp, err := l.IntegrationConsumeSettingDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
