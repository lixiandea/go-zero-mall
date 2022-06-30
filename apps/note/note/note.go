package main

import (
	"flag"
	"fmt"

	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/config"
	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/server"
	"github.com/lixiandea/go-zero-mall/apps/note/note/internal/svc"
	"github.com/lixiandea/go-zero-mall/apps/note/note/note"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/note.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewNoteServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		note.RegisterNoteServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
