package routes

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/valyala/fasthttp"

	"github.com/cosasdepuma/masterchef/backend/app/helpers"
)

type (
	dnsLookupParams struct {
		IPs []string `json:"ips"`
	}

	dnsLookupResponse struct {
		IPs      []string `json:"ips"`
		Domains  []string `json:"domains"`
		Warnings []string `json:"warnings"`
	}
)

// DNSLookup route that reverse lookup an IP address
func DNSLookup(ctx *fasthttp.RequestCtx, _ *fasthttp.Client, threads int) {
	// incoming request
	var params dnsLookupParams
	if json.Unmarshal(ctx.PostBody(), &params) != nil {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// content type
	ctx.SetContentType("application/json; charset=utf-8")
	// concurrency
	var warns []string = []string{}
	var domains []string = []string{}
	var lock sync.Mutex = sync.Mutex{}
	var pool chan struct{} = make(chan struct{}, threads)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(params.IPs))
	// iteration
	for _, ip := range params.IPs {
		pool <- struct{}{}
		go func(ip string) {
			// action
			addr, err := net.LookupAddr(ip)
			if err != nil {
				lock.Lock()
				warns = append(warns, err.Error())
				lock.Unlock()
			} else {
				// response
				lock.Lock()
				for _, domain := range addr {
					domains = append(domains, strings.TrimSuffix(domain, "."))
				}
				lock.Unlock()
			}
			<-pool
			wg.Done()
		}(ip)
	}
	wg.Wait()
	close(pool)
	// outgoing response
	var response dnsLookupResponse = dnsLookupResponse{
		IPs:      params.IPs,
		Domains:  helpers.UniqStrings(domains),
		Warnings: warns,
	}
	result, _ := json.Marshal(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "%s", result)
}
