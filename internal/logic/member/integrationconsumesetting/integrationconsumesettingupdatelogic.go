package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingUpdateLogic {
	return IntegrationConsumeSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingUpdateLogic) IntegrationConsumeSettingUpdate(req types.UpdateIntegrationConsumeSettingReq) (*types.UpdateIntegrationConsumeSettingResp, error) {
	_, err := l.svcCtx.UmsRpc.IntegrationConsumeSettingUpdate(l.ctx, &umsclient.IntegrationConsumeSettingUpdateReq{
		Id:                 req.Id,
		DeductionPerAmount: req.DeductionPerAmount,
		MaxPercentPerOrder: req.MaxPercentPerOrder,
		UseUnit:            req.UseUnit,
		CouponStatus:       req.CouponStatus,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
