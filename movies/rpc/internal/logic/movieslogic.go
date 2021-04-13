package logic

import (
	"context"

	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"

	"github.com/tal-tech/go-zero/core/logx"
)

type MoviesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoviesLogic {
	return &MoviesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MoviesLogic) Movies(in *movie.MovieReq) (*movie.MovieListResp, error) {
	// todo: add your logic here and delete this line

	return &movie.MovieListResp{}, nil
}
