package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderDeleteLogic {
	return OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req types.DeleteOrderReq) (*types.DeleteOrderResp, error) {
	// todo: add your logic here and delete this line
	_, _ = l.svcCtx.OmsRpc.OrderDelete(l.ctx, &omsclient.OrderDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteOrderResp{}, nil
}
