package logic

import (
	"context"

	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *movie.MovieRequest) (*movie.MovieResponse, error) {
	// todo: add your logic here and delete this line

	return &movie.MovieResponse{}, nil
}
