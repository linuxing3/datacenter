package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandListLogic {
	return ProductBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductBrandListLogic) ProductBrandList(req types.ListProductBrandReq) (*types.ListProductBrandResp, error) {
	resp, err := l.svcCtx.PmsRpc.BrandList(l.ctx, &pmsclient.BrandListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtProductBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListtProductBrandData{
			Id:                  item.Id,
			Name:                item.Name,
			FirstLetter:         item.FirstLetter,
			Sort:                item.Sort,
			FactoryStatus:       item.FactoryStatus,
			ShowStatus:          item.ShowStatus,
			ProductCount:        item.ProductCount,
			ProductCommentCount: item.ProductCommentCount,
			Logo:                item.Logo,
			BigPic:              item.BigPic,
			BrandStory:          item.BrandStory,
		})
	}

	return &types.ListProductBrandResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
