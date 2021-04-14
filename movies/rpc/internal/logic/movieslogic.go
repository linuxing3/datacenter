package logic

import (
	"context"
	"log"

	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"
	"datacenter/shared"

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

func (l *MoviesLogic) Movies(in *movie.MovieListReq) (*movie.MovieListResp, error) {
	// todo: add your logic here and delete this line
	movieInt := shared.StrToInt64(in.Ids[0])

	log.Printf("%v", movieInt)

	movieInfo, _ := l.svcCtx.MoviesModel.FindOne(movieInt)
	list := make([]*movie.MovieResp, 0)
	list = append(list, &movie.MovieResp{
		Title: movieInfo.Title,
		Description: movieInfo.Description,
		Url: movieInfo.Url,
	})
	return &movie.MovieListResp{
		Data: list,
	}, nil
	// return &movie.MovieListResp{}, nil
}
