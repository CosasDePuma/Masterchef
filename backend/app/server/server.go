package server

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/cosasdepuma/misterchef/backend/app/config"
)

// Start the backend server
func Start(conf *config.Config) {
	fmt.Printf(" ** Server started! **\nhttp://%s/\n", conf.Server.Address)
	var router func(*fasthttp.RequestCtx) = newRouter(conf)
	var server *fasthttp.Server = &fasthttp.Server{
		Handler:           router,
		DisableKeepalive:  true,
		CloseOnShutdown:   true,
		StreamRequestBody: true,
	}
	if err := server.ListenAndServe(conf.Server.Address); err != nil {
		fmt.Printf(" ** ! %s\n", err)
	}
}
