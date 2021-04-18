package logic

import (
	"context"
	"fmt"

	"datacenter/internal/svc"
	"datacenter/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req types.ListOrderReq) (*types.ListOrderResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.OmsRpc.OrderList(l.ctx, &omsclient.OrderListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtOrderData

	for _, item := range resp.List {
		list = append(list, &types.ListtOrderData{
			Id:                    item.Id,
			MemberId:              item.MemberId,
			CouponId:              item.CouponId,
			OrderSn:               item.OrderSn,
			CreateTime:            item.CreateTime,
			MemberUsername:        item.MemberUsername,
			TotalAmount:           item.TotalAmount,
			PayAmount:             item.PayAmount,
			FreightAmount:         item.FreightAmount,
			PromotionAmount:       item.PromotionAmount,
			IntegrationAmount:     item.IntegrationAmount,
			CouponAmount:          item.CouponAmount,
			DiscountAmount:        item.DiscountAmount,
			PayType:               item.PayType,
			SourceType:            item.SourceType,
			Status:                item.Status,
			OrderType:             item.OrderType,
			DeliveryCompany:       item.DeliveryCompany,
			DeliverySn:            item.DeliverySn,
			AutoConfirmDay:        item.AutoConfirmDay,
			Integration:           item.Integration,
			Growth:                item.Growth,
			PromotionInfo:         item.PromotionInfo,
			BillType:              item.BillType,
			BillHeader:            item.BillHeader,
			BillContent:           item.BillContent,
			BillReceiverPhone:     item.BillReceiverPhone,
			BillReceiverEmail:     item.BillReceiverEmail,
			ReceiverName:          item.ReceiverName,
			ReceiverPhone:         item.ReceiverPhone,
			ReceiverPostCode:      item.ReceiverPostCode,
			ReceiverProvince:      item.ReceiverProvince,
			ReceiverCity:          item.ReceiverCity,
			ReceiverRegion:        item.ReceiverRegion,
			ReceiverDetailAddress: item.ReceiverDetailAddress,
			Note:                  item.Note,
			ConfirmStatus:         item.ConfirmStatus,
			DeleteStatus:          item.DeleteStatus,
			UseIntegration:        item.UseIntegration,
			PaymentTime:           item.PaymentTime,
			DeliveryTime:          item.DeliveryTime,
			ReceiveTime:           item.ReceiveTime,
			CommentTime:           item.CommentTime,
			ModifyTime:            item.ModifyTime,
		})
	}
	return &types.ListOrderResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
