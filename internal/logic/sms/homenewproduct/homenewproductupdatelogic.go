package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeNewProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductUpdateLogic {
	return HomeNewProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductUpdateLogic) HomeNewProductUpdate(req types.UpdateHomeNewProductReq) (*types.UpdateHomeNewProductResp, error) {
	_, err := l.svcCtx.SmsRpc.HomeNewProductUpdate(l.ctx, &smsclient.HomeNewProductUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateHomeNewProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
