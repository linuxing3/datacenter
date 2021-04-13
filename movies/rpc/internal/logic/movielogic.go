package logic

import (
	"context"

	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovieLogic {
	return &MovieLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MovieLogic) Movie(in *movie.MovieListReq) (*movie.MovieResp, error) {
	// todo: add your logic here and delete this line

	return &movie.MovieResp{}, nil
}
