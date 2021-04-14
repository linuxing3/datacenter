package logic

import (
	"context"
	"fmt"
	"log"

	"datacenter/internal/svc"
	"datacenter/internal/types"
	// FIXED: 调用rpcclient
	"datacenter/movies/rpc/movieclient"
	"github.com/tal-tech/go-zero/core/logx"
)

type MovieInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovieInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) MovieInfoLogic {
	return MovieInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovieInfoLogic) MovieInfo(req types.MovieReq) (*movieclient.MovieListResp, error) {
	// FIXED: 调用后端MoviesRpc服务器，发送请求
	// 检查 movies/rpc/internal/logic/movieslogic.go
	id:= fmt.Sprintf("%s", req.Id)
	log.Printf("从请求中获取id： %v", id);

	ids := []string{id}
	return l.svcCtx.MoviesRpc.Movies(l.ctx, &movieclient.MovieListReq{
		Ids: ids,
	})
}
