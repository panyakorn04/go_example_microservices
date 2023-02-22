package logic

import (
	"context"
	"go-zero-demo/posts/internal/svc"
	"go-zero-demo/posts/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsLogic {
	return &PostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsLogic) Posts() (*types.NewPost, error) {
	return &types.NewPost{
		Title:       "Hello, World!",
		Description: "This is a post",
	}, nil
}
