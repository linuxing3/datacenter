package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateDeleteLogic {
	return FeightTemplateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateDeleteLogic) FeightTemplateDelete(req types.DeleteFeightTemplateReq) (*types.DeleteFeightTemplateResp, error) {
	_, _ = l.svcCtx.PmsRpc.FeightTemplateDelete(l.ctx, &pmsclient.FeightTemplateDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteFeightTemplateResp{
		Code:    "000000",
		Message: "",
	}, nil
}
