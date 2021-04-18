package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/oms/omsclient"
	"time"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OperateHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OperateHistoryListLogic {
	return OperateHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHistoryListLogic) OperateHistoryList(req types.ListOperateHistoryReq) (*types.ListOperateHistoryResp, error) {
	resp, err := l.svcCtx.OmsRpc.OrderOperateHistoryList(l.ctx, &omsclient.OrderOperateHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtOperateHistoryData

	for _, item := range resp.List {
		CreateTime, _ := time.Parse("2006-01-02 15:04:05", item.CreateTime)
		list = append(list, &types.ListtOperateHistoryData{
			Id:          item.Id,
			OrderId:     item.OrderId,
			OperateMan:  item.OperateMan,
			CreateTime:  CreateTime,
			OrderStatus: item.OrderStatus,
			Note:        item.Note,
		})
	}

	return &types.ListOperateHistoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
