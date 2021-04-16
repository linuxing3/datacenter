package logic

import (
	"context"

	adder "datacenter/bookstore/rpc/adder/adderclient"
	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	// t: add your logic here and delete this line
	resp, err := l.svcCtx.AdderRpc.Add(l.ctx, &adder.AddReq{
		Book: req.Book,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: resp.Ok,
	}, nil
}
