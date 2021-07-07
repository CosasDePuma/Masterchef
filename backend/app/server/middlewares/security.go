// https://wiki.owasp.org/index.php/OWASP_Secure_Headers_Project#tab=Headers

package middlewares

import "github.com/valyala/fasthttp"

// Security adds response headers related to server security
func Security(ctx *fasthttp.RequestCtx) {
	// ctx.Response.Header.Set("Content-Security-Policy", "default-src 'self'") // FIXME: CSP                                                                  // https://developer.mozilla.org/es/docs/Web/HTTP/CSP
	ctx.Response.Header.Set("Expect-CT", "max-age=86400, enforce")                                                                                // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT
	ctx.Response.Header.Set("Feature-Policy", "camera 'none'; geolocation 'none'; microphone 'none'; payment 'none'; usb 'none'; vibrate 'none'") // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Feature-Policy
	ctx.Response.Header.Set("Referer", "no-referer")                                                                                              // https://developer.mozilla.org/es/docs/Web/HTTP/Headers/Referrer-Policy
	ctx.Response.Header.Set("Referrer-Policy", "no-referrer")                                                                                     // https://developer.mozilla.org/es/docs/Web/HTTP/Headers/Referrer-Policy
	ctx.Response.Header.Set("Strict-Transport-Security", "max-age=31536000")                                                                      // https://developer.mozilla.org/es/docs/Web/HTTP/Headers/Strict-Transport-Security
	ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")                                                                                  // https://developer.mozilla.org/es/docs/Web/HTTP/Headers/X-Content-Type-Options
	ctx.Response.Header.Set("X-Frame-Options", "deny")                                                                                            // https://developer.mozilla.org/es/docs/Web/HTTP/Headers/X-Frame-Options
	ctx.Response.Header.Set("X-XSS-Protection", "1; mode=block")                                                                                  // https://developer.mozilla.org/es/docs/Web/HTTP/Headers/X-XSS-Protection
	ctx.Response.Header.Set("X-Permitted-Cross-Domain-Policies", "none")                                                                          // https://wiki.owasp.org/index.php/OWASP_Secure_Headers_Project#xpcdp
}
