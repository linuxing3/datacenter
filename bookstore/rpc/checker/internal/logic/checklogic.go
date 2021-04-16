package logic

import (
	"context"

	"datacenter/bookstore/rpc/checker/check"
	"datacenter/bookstore/rpc/checker/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.BookModel.FindOne(in.Book)
	if err != nil {
		return nil, err
	}
	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
}
