package logic

import (
	"context"
	"log"

	"datacenter/bookstore/model"

	"datacenter/bookstore/rpc/adder/add"
	"datacenter/bookstore/rpc/adder/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	log.Printf("in object : %v",in)
	_, err := l.svcCtx.BookModel.Insert(model.Book{
		Book: in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}
	return &add.AddResp{
		Ok: true,
	}, nil
}
