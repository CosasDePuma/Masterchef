package server

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"cosasdepuma/masterchef/app/config"
	"cosasdepuma/masterchef/app/server/client"
	"cosasdepuma/masterchef/app/server/middlewares"
	"cosasdepuma/masterchef/app/server/routes"
	"cosasdepuma/masterchef/public"
)

// API indicates the path and version relative to them
const API = "/api/v1"

var (
	clnt                                                                 = client.New()
	sitemap map[string]func(*fasthttp.RequestCtx, *fasthttp.Client, int) = map[string]func(*fasthttp.RequestCtx, *fasthttp.Client, int){
		"/dummy":               routes.Dummy,
		"/request":             routes.Request,
		"/lookup/dns":          routes.DNSLookup,
		"/lookup/ip":           routes.IPLookup,
		"/stealth/spider":      routes.Spider,
		"/stealth/subdomainer": routes.Subdomainer,
	}
)

func newRouter(conf *config.Config) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// pre-middlewares
		middlewares.CORS(ctx)
		middlewares.Meta(ctx)
		middlewares.Security(ctx)
		var path string = string(ctx.Path())
		// methods
		switch string(ctx.Request.Header.Method()) {
		case fasthttp.MethodOptions:
			// cors
			ctx.SetStatusCode(fasthttp.StatusOK)
		case fasthttp.MethodGet:
			// frontend
			ctx.SetStatusCode(fasthttp.StatusOK)
			switch path {
			case "/", "/index.html":
				ctx.SetContentType("text/html; charset=utf-8")
				fmt.Fprint(ctx, public.Index)
			case "/app.js":
				ctx.SetContentType("text/javascript; charset=utf-8")
				fmt.Fprint(ctx, public.App)
			default:
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			}
		case fasthttp.MethodPost:
			// check path & user-agent
			if !strings.HasPrefix(path, API) || !bytes.EqualFold(ctx.Request.Header.Peek("X-Powered-By"), []byte("Masterchef!")) {
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			} else {
				// api
				if endpoint, ok := sitemap[strings.TrimPrefix(path, API)]; !ok {
					ctx.SetStatusCode(fasthttp.StatusNotFound)
				} else {
					endpoint(ctx, clnt, conf.Threads)
				}
			}
		}
		// post-middlewares
		middlewares.Logger(ctx)
	}
}
