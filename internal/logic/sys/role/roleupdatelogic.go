package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleUpdateLogic {
	return RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req types.UpdateRoleReq) (*types.UpdateRoleResp, error) {
	_, err := l.svcCtx.SysRpc.RoleUpdate(l.ctx, &sysclient.RoleUpdateReq{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
		//todo 从token里面拿
		LastUpdateBy: "admin",
		Status:       req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateRoleResp{
		Code:    "000000",
		Message: "",
	}, nil
}
