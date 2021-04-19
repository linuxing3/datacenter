package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigAddLogic {
	return ConfigAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigAddLogic) ConfigAdd(req types.AddConfigReq) (*types.AddConfigResp, error) {
	_, err := l.svcCtx.SysRpc.ConfigAdd(l.ctx, &sysclient.ConfigAddReq{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        req.Sort,
		Remarks:     req.Remarks,
		//todo 从token里面拿
		CreateBy: "admin",
	})

	if err != nil {
		return nil, err
	}

	return &types.AddConfigResp{}, nil
}
