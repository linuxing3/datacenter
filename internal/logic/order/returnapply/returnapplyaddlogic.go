package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnApplyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyAddLogic {
	return ReturnApplyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyAddLogic) ReturnApplyAdd(req types.AddReturnApplyReq) (*types.AddReturnApplyResp, error) {
	_, err := l.svcCtx.OmsRpc.OrderReturnApplyAdd(l.ctx, &omsclient.OrderReturnApplyAddReq{
		OrderId:          req.OrderId,
		CompanyAddressId: req.CompanyAddressId,
		ProductId:        req.ProductId,
		OrderSn:          req.OrderSn,
		CreateTime:       req.CreateTime,
		MemberUsername:   req.MemberUsername,
		ReturnAmount:     int64(req.ReturnAmount),
		ReturnName:       req.ReturnName,
		ReturnPhone:      req.ReturnPhone,
		Status:           req.Status,
		HandleTime:       req.HandleTime,
		ProductPic:       req.ProductPic,
		ProductName:      req.ProductName,
		ProductBrand:     req.ProductBrand,
		ProductAttr:      req.ProductAttr,
		ProductCount:     req.ProductCount,
		ProductPrice:     int64(req.ProductPrice),
		ProductRealPrice: int64(req.ProductRealPrice),
		Reason:           req.Reason,
		Description:      req.Description,
		ProofPics:        req.ProofPics,
		HandleNote:       req.HandleNote,
		HandleMan:        req.HandleMan,
		ReceiveMan:       req.ReceiveMan,
		ReceiveTime:      req.ReceiveTime,
		ReceiveNote:      req.ReceiveNote,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddReturnApplyResp{
		Code:    "000000",
		Message: "",
	}, nil
}
