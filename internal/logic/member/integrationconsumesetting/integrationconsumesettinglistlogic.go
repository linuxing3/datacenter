package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/ums/umsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingListLogic {
	return IntegrationConsumeSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingListLogic) IntegrationConsumeSettingList(req types.ListIntegrationConsumeSettingReq) (*types.ListIntegrationConsumeSettingResp, error) {
	resp, err := l.svcCtx.UmsRpc.IntegrationConsumeSettingList(l.ctx, &umsclient.IntegrationConsumeSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtIntegrationConsumeSettingData

	for _, item := range resp.List {
		list = append(list, &types.ListtIntegrationConsumeSettingData{
			Id:                 item.Id,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
		})
	}

	return &types.ListIntegrationConsumeSettingResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
