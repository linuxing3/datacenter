package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"
	"strconv"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictAddLogic {
	return DictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictAddLogic) DictAdd(req types.AddDictReq) (*types.AddDictResp, error) {
	sort, _ := strconv.ParseInt(req.Sort, 10, 64)
	_, err := l.svcCtx.SysRpc.DictAdd(l.ctx, &sysclient.DictAddReq{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        sort,
		Remarks:     req.Remarks,
		//todo 从token里面拿
		CreateBy: "admin",
	})

	if err != nil {
		return nil, err
	}

	return &types.AddDictResp{
		Code:    "000000",
		Message: "",
	}, nil
}
