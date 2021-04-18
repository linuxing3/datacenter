package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingDeleteLogic {
	return IntegrationConsumeSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(req types.DeleteIntegrationConsumeSettingReq) (*types.DeleteIntegrationConsumeSettingResp, error) {
	_, _ = l.svcCtx.UmsRpc.IntegrationConsumeSettingDelete(l.ctx, &umsclient.IntegrationConsumeSettingDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
