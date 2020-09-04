package server

import (
	http_render "github.com/moisespsena-go/http-render"
	"github.com/moisespsena-go/http-render/ropt"

	"github.com/moisespsena-go/assetfs/assetfsapi"
	"github.com/moisespsena-go/httpu"
	"github.com/moisespsena-go/task"
	"github.com/moisespsena-go/xroute"
)

type Server struct {
	srv *httpu.Server
}

type Config struct {
	httpu.Config
	FS assetfsapi.Interface
}

func New(cfg *Config) task.Task {
	if len(cfg.Listeners) == 0 {
		cfg.Listeners = append(cfg.Listeners, httpu.ListenerConfig{
			Addr: "0.0.0.0:12000",
		})
	}

	router := xroute.NewMux()

	Render := http_render.New(ropt.FS(cfg.FS.NameSpace("templates")))

	router.Mount("/static", cfg.FS.NameSpace("static"))
	router.Mount("/", Render)

	return httpu.NewServer(&cfg.Config, router)
}
