package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SysLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SysLogDeleteLogic {
	return SysLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogDeleteLogic) SysLogDelete(req types.DeleteSysLogReq) (*types.DeleteSysLogResp, error) {
	_, err := l.svcCtx.SysRpc.SysLogDelete(l.ctx, &sysclient.SysLogDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.DeleteSysLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
