package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryAddLogic {
	return OperateHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryAddLogic) OperateHistoryAdd(req types.AddOperateHistoryReq) (*types.AddOperateHistoryResp, error) {
	_, err := l.svcCtx.OmsRpc.OrderOperateHistoryAdd(l.ctx, &omsclient.OrderOperateHistoryAddReq{
		OrderId:     req.OrderId,
		OperateMan:  req.OperateMan,
		CreateTime:  req.CreateTime.Format("2006-01-02 15:04:05"),
		OrderStatus: req.OrderStatus,
		Note:        req.Note,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddOperateHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
