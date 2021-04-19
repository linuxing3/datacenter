package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptAddLogic {
	return DeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptAddLogic) DeptAdd(req types.AddDeptReq) (*types.AddDeptResp, error) {
	_, err := l.svcCtx.SysRpc.DeptAdd(l.ctx, &sysclient.DeptAddReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		OrderNum: req.OrderNum,
		//todo 从token里面拿
		CreateBy: "admin",
	})

	if err != nil {
		return nil, err
	}

	return &types.AddDeptResp{
		Code:    "000000",
		Message: "",
	}, nil
}
