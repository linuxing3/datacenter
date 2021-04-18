package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponAddLogic {
	return CouponAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponAddLogic) CouponAdd(req types.AddCouponReq) (*types.AddCouponResp, error) {
	_, err := l.svcCtx.SmsRpc.CouponAdd(l.ctx, &smsclient.CouponAddReq{
		Type:         req.Type,
		Name:         req.Name,
		Platform:     req.Platform,
		Count:        req.Count,
		Amount:       req.Amount,
		PerLimit:     req.PerLimit,
		MinPoint:     req.MinPoint,
		StartTime:    req.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:      req.EndTime.Format("2006-01-02 15:04:05"),
		UseType:      req.UseType,
		Note:         req.Note,
		PublishCount: req.PublishCount,
		UseCount:     req.UseCount,
		ReceiveCount: req.ReceiveCount,
		EnableTime:   req.EnableTime,
		Code:         req.Code,
		MemberLevel:  req.MemberLevel,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddCouponResp{
		Code:    "000000",
		Message: "",
	}, nil
}
