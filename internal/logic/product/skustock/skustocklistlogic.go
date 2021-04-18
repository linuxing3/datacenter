package logic

import (
	"context"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"fmt"
	"go-zero-admin/rpc/pms/pmsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type SkuStockListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SkuStockListLogic {
	return SkuStockListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuStockListLogic) SkuStockList(req types.ListSkuStockReq) (*types.ListSkuStockResp, error) {
	resp, err := l.svcCtx.PmsRpc.SkuStockList(l.ctx, &pmsclient.SkuStockListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtSkuStockData

	for _, item := range resp.List {
		list = append(list, &types.ListtSkuStockData{
			Id:             item.Id,
			ProductId:      item.ProductId,
			SkuCode:        item.SkuCode,
			Price:          float64(item.Price),
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: float64(item.PromotionPrice),
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	return &types.ListSkuStockResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
