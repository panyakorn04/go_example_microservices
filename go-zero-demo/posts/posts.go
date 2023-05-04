package main

import (
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/example/graceful/dns/api/handler"
	"github.com/tal-tech/go-zero/rest"
	"honnef.co/go/tools/config"
)

var configFile = flag.String("f", "etc/posts-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
