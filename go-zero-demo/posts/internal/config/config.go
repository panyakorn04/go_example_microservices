package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// Add you own custom config fields here
	DataSource string
	CacheRedis cache.CacheConf // redis cache
}
