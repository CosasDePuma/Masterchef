package middlewares

import "github.com/valyala/fasthttp"

// CORS adds response headers related to Cross-Origin Resource Sharing
func CORS(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, X-Powered-By")
	ctx.Response.Header.Set("Access-Control-Expose-Headers", "Content-Type, Server")
}
