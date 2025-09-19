package main

import (
	context "context"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/pkg/printx"
	_ "go.uber.org/automaxprocs"

	"github.com/lucyandlucky/golang-platform-configure/internal/app"
	"github.com/lucyandlucky/golang-platform-configure/internal/conf"
)

const (
	AppName = "Configure"
)

func main() {
	srv := kratosx.New(
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(kratos.AfterStart(func(_ context.Context) error {
			printx.ArtFont(fmt.Sprintf("Hello %s", AppName))

			return nil
		})),
	)

	if err := srv.Run(); err != nil {
		log.Fatal("run service fail", err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	cfg := &conf.Config{}
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Error("business 配置变更失败")
		} else {
			log.Error("file 配置变更成功")
		}
	})

	app.Register(cfg, hs, gs)
}
