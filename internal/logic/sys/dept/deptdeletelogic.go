package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptDeleteLogic {
	return DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptDeleteLogic) DeptDelete(req types.DeleteDeptReq) (*types.DeleteDeptResp, error) {
	_, err := l.svcCtx.SysRpc.DeptDelete(l.ctx, &sysclient.DeptDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteDeptResp{
		Code:    "000000",
		Message: "",
	}, nil
}
