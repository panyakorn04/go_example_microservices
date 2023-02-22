package svc

import (
	"go-zero-demo/posts/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// Add you own custom service context fields here
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
