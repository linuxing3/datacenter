package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionAddLogic {
	return FlashPromotionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionAddLogic) FlashPromotionAdd(req types.AddFlashPromotionReq) (*types.AddFlashPromotionResp, error) {
	_, err := l.svcCtx.SmsRpc.FlashPromotionAdd(l.ctx, &smsclient.FlashPromotionAddReq{
		Title:      req.Title,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Status:     req.Status,
		CreateTime: req.CreateTime,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddFlashPromotionResp{
		Code:    "000000",
		Message: "",
	}, nil
}
