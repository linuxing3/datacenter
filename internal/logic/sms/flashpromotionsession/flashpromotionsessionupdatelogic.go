package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionSessionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionUpdateLogic {
	return FlashPromotionSessionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionUpdateLogic) FlashPromotionSessionUpdate(req types.UpdateFlashPromotionSessionReq) (*types.UpdateFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.SmsRpc.FlashPromotionSessionUpdate(l.ctx, &smsclient.FlashPromotionSessionUpdateReq{
		Id:         req.Id,
		Name:       req.Name,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateFlashPromotionSessionResp{
		Code:    "000000",
		Message: "",
	}, nil
}
