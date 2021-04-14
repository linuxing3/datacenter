// Code generated by goctl. DO NOT EDIT!
// Source: movie.proto

package server

import (
	"context"

	"datacenter/movies/rpc/internal/logic"
	"datacenter/movies/rpc/internal/svc"
	"datacenter/movies/rpc/movie"
)

type MoviesServer struct {
	svcCtx *svc.ServiceContext
}

func NewMoviesServer(svcCtx *svc.ServiceContext) *MoviesServer {
	return &MoviesServer{
		svcCtx: svcCtx,
	}
}

func (s *MoviesServer) Movies(ctx context.Context, in *movie.MovieListReq) (*movie.MovieListResp, error) {
	l := logic.NewMoviesLogic(ctx, s.svcCtx)
	return l.Movies(in)
}

func (s *MoviesServer) Movie(ctx context.Context, in *movie.MovieReq) (*movie.MovieResp, error) {
	l := logic.NewMovieLogic(ctx, s.svcCtx)
	return l.Movie(in)
}
