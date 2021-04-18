package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryListLogic {
	return CouponHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryListLogic) CouponHistoryList(req types.ListCouponHistoryReq) (*types.ListCouponHistoryResp, error) {
	resp, err := l.svcCtx.SmsRpc.CouponHistoryList(l.ctx, &smsclient.CouponHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtCouponHistoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtCouponHistoryData{
			Id:             item.Id,
			CouponId:       item.CouponId,
			MemberId:       item.MemberId,
			CouponCode:     item.CouponCode,
			MemberNickname: item.MemberNickname,
			GetType:        item.GetType,
			CreateTime:     item.CreateTime,
			UseStatus:      item.UseStatus,
			UseTime:        item.UseTime,
			OrderId:        item.OrderId,
			OrderSn:        item.OrderSn,
		})
	}

	return &types.ListCouponHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
