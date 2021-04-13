package handler

import (
	"net/http"

	"datacenter/movies/api/internal/logic/movie"
	"datacenter/movies/api/internal/svc"
	"datacenter/movies/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MovieInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MovieReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMovieInfoLogic(r.Context(), ctx)
		resp, err := l.MovieInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
