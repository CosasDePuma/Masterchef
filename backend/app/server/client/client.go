package client

import "github.com/valyala/fasthttp"

// New client with default settings
func New() *fasthttp.Client {
	return &fasthttp.Client{
		Name:                     "Masterchef!",
		NoDefaultUserAgentHeader: true,
	}
}
