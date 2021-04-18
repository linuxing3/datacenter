package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandDeleteLogic {
	return ProductBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandDeleteLogic) ProductBrandDelete(req types.DeleteProductBrandReq) (*types.DeleteProductBrandResp, error) {
	_, _ = l.svcCtx.PmsRpc.BrandDelete(l.ctx, &pmsclient.BrandDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteProductBrandResp{
		Code:    "000000",
		Message: "",
	}, nil
}
