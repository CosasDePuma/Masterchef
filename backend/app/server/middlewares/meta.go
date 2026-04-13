package middlewares

import "github.com/valyala/fasthttp"

// Meta adds the response headers relative to the server
func Meta(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Server", "Masterchef")
}
