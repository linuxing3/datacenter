package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SkuStockDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockDeleteLogic {
	return SkuStockDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockDeleteLogic) SkuStockDelete(req types.DeleteSkuStockReq) (*types.DeleteSkuStockResp, error) {
	_, _ = l.svcCtx.PmsRpc.SkuStockDelete(l.ctx, &pmsclient.SkuStockDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteSkuStockResp{
		Code:    "000000",
		Message: "",
	}, nil
}
