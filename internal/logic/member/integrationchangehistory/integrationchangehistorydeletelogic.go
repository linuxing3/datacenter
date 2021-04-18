package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationChangeHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationChangeHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationChangeHistoryDeleteLogic {
	return IntegrationChangeHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationChangeHistoryDeleteLogic) IntegrationChangeHistoryDelete(req types.DeleteIntegrationChangeHistoryReq) (*types.DeleteIntegrationChangeHistoryResp, error) {
	_, _ = l.svcCtx.UmsRpc.IntegrationChangeHistoryDelete(l.ctx, &umsclient.IntegrationChangeHistoryDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteIntegrationChangeHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
