package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/posts/internal/logic"
	"net/http"

	"go-zero-demo/posts/internal/svc"
)

func PostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPostsLogic(r.Context(), svcCtx)
		resp, err := l.Posts()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
