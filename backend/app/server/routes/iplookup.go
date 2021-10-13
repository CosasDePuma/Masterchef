package routes

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/valyala/fasthttp"

	"github.com/cosasdepuma/misterchef/backend/app/helpers"
)

type (
	ipLookupParams struct {
		Domains []string `json:"domains"`
		OnlyV4  bool     `json:"onlyv4"`
	}

	ipLookupResponse struct {
		Domains  []string `json:"domains"`
		IPs      []string `json:"ips"`
		Warnings []string `json:"warnings"`
	}
)

// IPLookup route that get the IP addresses from a domain
func IPLookup(ctx *fasthttp.RequestCtx, _ *fasthttp.Client, threads int) {
	// incoming request
	var params ipLookupParams
	if json.Unmarshal(ctx.PostBody(), &params) != nil {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// content type
	ctx.SetContentType("application/json; charset=utf-8")
	// concurrency
	var warns []string = []string{}
	var ips []string = []string{}
	var lock sync.Mutex = sync.Mutex{}
	var pool chan struct{} = make(chan struct{}, threads)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(params.Domains))
	// iteration
	for _, domain := range params.Domains {
		pool <- struct{}{}
		go func(domain string) {
			// action
			addr, err := net.LookupIP(domain)
			if err != nil {
				lock.Lock()
				warns = append(warns, err.Error())
				lock.Unlock()
			} else {
				// response
				lock.Lock()
				if params.OnlyV4 {
					for _, ip := range addr {
						if ipv4 := ip.To4(); ipv4 != nil {
							ips = append(ips, ipv4.String())
						}
					}
				} else {
					for _, ip := range addr {
						ips = append(ips, ip.String())
					}
				}
				lock.Unlock()
			}
			<-pool
			wg.Done()
		}(domain)
	}
	wg.Wait()
	close(pool)
	// outgoing response
	var response ipLookupResponse = ipLookupResponse{
		Domains:  params.Domains,
		IPs:      helpers.UniqStrings(ips),
		Warnings: warns,
	}
	result, _ := json.Marshal(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "%s", result)
}
