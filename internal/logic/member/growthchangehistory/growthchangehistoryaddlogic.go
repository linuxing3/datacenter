package logic

import (
	"context"
	"datacenter/internal/svc"
	"datacenter/internal/types"
	"go-zero-admin/rpc/ums/umsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GrowthChangeHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrowthChangeHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) GrowthChangeHistoryAddLogic {
	return GrowthChangeHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrowthChangeHistoryAddLogic) GrowthChangeHistoryAdd(req types.AddGrowthChangeHistoryReq) (*types.AddGrowthChangeHistoryResp, error) {
	_, err := l.svcCtx.UmsRpc.GrowthChangeHistoryAdd(l.ctx, &umsclient.GrowthChangeHistoryAddReq{
		MemberId:    req.MemberId,
		CreateTime:  req.CreateTime,
		ChangeType:  req.ChangeType,
		ChangeCount: req.ChangeCount,
		OperateMan:  req.OperateMan,
		OperateNote: req.OperateNote,
		SourceType:  req.SourceType,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddGrowthChangeHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
