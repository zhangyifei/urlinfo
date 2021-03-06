package main

import (
	"flag"
	"fmt"

	"urlinfo/urlinfo/internal/config"
	"urlinfo/urlinfo/internal/handler"
	"urlinfo/urlinfo/internal/svc"
	urlinfoRouter "urlinfo/urlinfo/pkg/router"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/urlinfo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)

	urlinfoRouter := urlinfoRouter.NewRouter()

	rest.WithRouter(urlinfoRouter)(server)

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
