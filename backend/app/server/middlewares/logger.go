package middlewares

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// Logger records the requests along with their response status
func Logger(ctx *fasthttp.RequestCtx) {
	fmt.Printf("%s | %s/%s | [%d] %s: %s\n", time.Now().Format("2006-01-02T15:04:05"), ctx.RemoteAddr().String(), ctx.RemoteAddr().Network(), ctx.Response.StatusCode(), ctx.Request.Header.Method(), ctx.Request.RequestURI())
}
