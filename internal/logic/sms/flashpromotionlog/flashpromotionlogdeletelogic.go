package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogDeleteLogic {
	return FlashPromotionLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogDeleteLogic) FlashPromotionLogDelete(req types.DeleteFlashPromotionLogReq) (*types.DeleteFlashPromotionLogResp, error) {
	_, _ = l.svcCtx.SmsRpc.FlashPromotionLogDelete(l.ctx, &smsclient.FlashPromotionLogDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteFlashPromotionLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
