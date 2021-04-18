package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCommentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentUpdateLogic {
	return ProductCommentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentUpdateLogic) ProductCommentUpdate(req types.UpdateProductCommentReq) (*types.UpdateProductCommentResp, error) {
	_, err := l.svcCtx.PmsRpc.CommentUpdate(l.ctx, &pmsclient.CommentUpdateReq{
		Id:               req.Id,
		ProductId:        req.ProductId,
		MemberNickName:   req.MemberNickName,
		ProductName:      req.ProductName,
		Star:             req.Star,
		MemberIp:         req.MemberIp,
		CreateTime:       req.CreateTime,
		ShowStatus:       req.ShowStatus,
		ProductAttribute: req.ProductAttribute,
		CollectCouont:    req.CollectCouont,
		ReadCount:        req.ReadCount,
		Content:          req.Content,
		Pics:             req.Pics,
		MemberIcon:       req.MemberIcon,
		ReplayCount:      req.ReplayCount,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateProductCommentResp{
		Code:    "000000",
		Message: "",
	}, nil
}
