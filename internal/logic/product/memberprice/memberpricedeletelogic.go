package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceDeleteLogic {
	return MemberPriceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceDeleteLogic) MemberPriceDelete(req types.DeleteMemberPriceReq) (*types.DeleteMemberPriceResp, error) {
	_, _ = l.svcCtx.PmsRpc.MemberPriceDelete(l.ctx, &pmsclient.MemberPriceDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberPriceResp{
		Code:    "000000",
		Message: "",
	}, nil
}
