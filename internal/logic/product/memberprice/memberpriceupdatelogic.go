package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceUpdateLogic {
	return MemberPriceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceUpdateLogic) MemberPriceUpdate(req types.UpdateMemberPriceReq) (*types.UpdateMemberPriceResp, error) {
	_, err := l.svcCtx.PmsRpc.MemberPriceUpdate(l.ctx, &pmsclient.MemberPriceUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		MemberLevelId:   req.MemberLevelId,
		MemberPrice:     int64(req.MemberPrice),
		MemberLevelName: req.MemberLevelName,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberPriceResp{
		Code:    "000000",
		Message: "",
	}, nil
}
