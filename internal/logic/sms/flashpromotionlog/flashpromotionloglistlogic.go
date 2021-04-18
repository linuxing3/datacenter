package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogListLogic {
	return FlashPromotionLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogListLogic) FlashPromotionLogList(req types.ListFlashPromotionLogReq) (*types.ListFlashPromotionLogResp, error) {
	resp, err := l.svcCtx.SmsRpc.FlashPromotionLogList(l.ctx, &smsclient.FlashPromotionLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtFlashPromotionLogData

	for _, item := range resp.List {
		list = append(list, &types.ListtFlashPromotionLogData{
			Id:            item.Id,
			MemberId:      item.MemberId,
			ProductId:     item.ProductId,
			MemberPhone:   item.MemberPhone,
			ProductName:   item.ProductName,
			SubscribeTime: item.SubscribeTime,
			SendTime:      item.SendTime,
		})
	}

	return &types.ListFlashPromotionLogResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
