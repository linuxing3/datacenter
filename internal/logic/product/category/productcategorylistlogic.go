package logic

import (
	"context"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"fmt"
	"go-zero-admin/rpc/pms/pmsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryListLogic {
	return ProductCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryListLogic) ProductCategoryList(req types.ListProductCategoryReq) (*types.ListProductCategoryResp, error) {
	resp, err := l.svcCtx.PmsRpc.ProductCategoryList(l.ctx, &pmsclient.ProductCategoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtProductCategoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtProductCategoryData{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			//ProductUnit:  strconv.FormatInt( item.ProductCount,10)+item.ProductUnit,
			ProductUnit: item.ProductUnit,
			NavStatus:   item.NavStatus,
			ShowStatus:  item.ShowStatus,
			Sort:        item.Sort,
			Icon:        item.Icon,
			Keywords:    item.Keywords,
			Description: item.Description,
		})
	}

	return &types.ListProductCategoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
