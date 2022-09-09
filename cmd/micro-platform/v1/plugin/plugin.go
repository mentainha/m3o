package plugin

import (
	"net/http"
	"strings"

	"github.com/micro/micro/v3/plugin"
	"github.com/micro/micro/v3/service/api/handler"
	"github.com/micro/micro/v3/service/api/server"
	"github.com/micro/micro/v3/service/config"
	"github.com/urfave/cli/v2"
	"m3o.dev/cmd/micro-platform/v1"
)

var (
	p = new(v1Plugin)
)

type v1Plugin struct {
	v1 handler.Handler
}

func (p *v1Plugin) Init(ctx *cli.Context) error {
	var address string

	val, err := config.Get("micro.v1.address")
	if err == nil {
		address = val.String("v1.micro:8080")
	}

	p.v1 = v1.NewHandler(address)

	// register the v1 handler
	server.Register("v1", p.v1)

	return nil
}

func (p *v1Plugin) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do default handler processing
		if !strings.HasPrefix(r.URL.Path, "/v1/") {
			h.ServeHTTP(w, r)
			return
		}

		// pass through key serving
		if strings.HasPrefix(r.URL.Path, "/v1/api/keys/") {
			h.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(r.URL.Path, "/v1/api/apis/") {
			h.ServeHTTP(w, r)
			return
		}

		// otherwise serve the v1 handler
		p.v1.ServeHTTP(w, r)
	})
}

func Register() {
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("v1"),
		plugin.WithInit(p.Init),
		plugin.WithHandler(p.Handler),
	))
}
