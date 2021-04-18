package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FlashPromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionListLogic {
	return FlashPromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionListLogic) FlashPromotionList(req types.ListFlashPromotionReq) (*types.ListFlashPromotionResp, error) {
	resp, err := l.svcCtx.SmsRpc.FlashPromotionList(l.ctx, &smsclient.FlashPromotionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtFlashPromotionData

	for _, item := range resp.List {
		list = append(list, &types.ListtFlashPromotionData{
			Id:         item.Id,
			Title:      item.Title,
			StartDate:  item.StartDate,
			EndDate:    item.EndDate,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ListFlashPromotionResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
