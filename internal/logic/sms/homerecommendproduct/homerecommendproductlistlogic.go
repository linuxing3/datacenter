package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductListLogic {
	return HomeRecommendProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductListLogic) HomeRecommendProductList(req types.ListHomeRecommendProductReq) (*types.ListHomeRecommendProductResp, error) {
	resp, err := l.svcCtx.SmsRpc.HomeRecommendProductList(l.ctx, &smsclient.HomeRecommendProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtHomeRecommendProductData

	for _, item := range resp.List {
		list = append(list, &types.ListtHomeRecommendProductData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &types.ListHomeRecommendProductResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
