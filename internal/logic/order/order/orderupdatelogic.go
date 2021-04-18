package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"
	"time"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderUpdateLogic {
	return OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req types.UpdateOrderReq) (*types.UpdateOrderResp, error) {
	// todo: add your logic here and delete this line
	_, _ = l.svcCtx.OmsRpc.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		Id: req.Id,
	})
	_, err := l.svcCtx.OmsRpc.OrderUpdate(l.ctx, &omsclient.OrderUpdateReq{
		Id:                    req.Id,
		MemberId:              req.MemberId,
		CouponId:              req.CouponId,
		OrderSn:               req.OrderSn,
		CreateTime:            req.CreateTime,
		MemberUsername:        req.MemberUsername,
		TotalAmount:           req.TotalAmount,
		PayAmount:             req.PayAmount,
		FreightAmount:         req.FreightAmount,
		PromotionAmount:       req.PromotionAmount,
		IntegrationAmount:     req.IntegrationAmount,
		CouponAmount:          req.CouponAmount,
		DiscountAmount:        req.DiscountAmount,
		PayType:               req.PayType,
		SourceType:            req.SourceType,
		Status:                req.Status,
		OrderType:             req.OrderType,
		DeliveryCompany:       req.DeliveryCompany,
		DeliverySn:            req.DeliverySn,
		AutoConfirmDay:        req.AutoConfirmDay,
		Integration:           req.Integration,
		Growth:                req.Growth,
		PromotionInfo:         req.PromotionInfo,
		BillType:              req.BillType,
		BillHeader:            req.BillHeader,
		BillContent:           req.BillContent,
		BillReceiverPhone:     req.BillReceiverPhone,
		BillReceiverEmail:     req.BillReceiverEmail,
		ReceiverName:          req.ReceiverName,
		ReceiverPhone:         req.ReceiverPhone,
		ReceiverPostCode:      req.ReceiverPostCode,
		ReceiverProvince:      req.ReceiverProvince,
		ReceiverCity:          req.ReceiverCity,
		ReceiverRegion:        req.ReceiverRegion,
		ReceiverDetailAddress: req.ReceiverDetailAddress,
		Note:                  req.Note,
		ConfirmStatus:         req.ConfirmStatus,
		DeleteStatus:          req.DeleteStatus,
		UseIntegration:        req.UseIntegration,
		PaymentTime:           time.Now().Format("2006-01-02 15:04:05"),
		DeliveryTime:          time.Now().Format("2006-01-02 15:04:05"),
		ReceiveTime:           time.Now().Format("2006-01-02 15:04:05"),
		CommentTime:           time.Now().Format("2006-01-02 15:04:05"),
		ModifyTime:            time.Now().Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return nil, err
	}
	return &types.UpdateOrderResp{}, nil
}
