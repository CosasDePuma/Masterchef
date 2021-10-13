package routes

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/valyala/fasthttp"

	"github.com/cosasdepuma/misterchef/backend/app/helpers"
	"github.com/cosasdepuma/misterchef/backend/app/providers/subdomainers"
)

type (
	subdomainerProvider struct {
		Endpoint func(*fasthttp.Client) string
		Get      func([]byte) []string
	}
	subdomainerParams struct {
		Domains  []string `json:"domains"`
		Provider string   `json:"provider"`
	}
	subdomainerResponse struct {
		Domains    []string `json:"domains"`
		Subdomains []string `json:"subdomains"`
		Warnings   []string `json:"warnings"`
	}
)

var subdomainerProviders map[string]subdomainerProvider = map[string]subdomainerProvider{
	"hackertarget": {
		Endpoint: subdomainers.HackerTargetEndpoint,
		Get:      subdomainers.HackerTargetResult,
	},
	"omnisint": {
		Endpoint: subdomainers.OmnisintEndpoint,
		Get:      subdomainers.OmnisintResult,
	},
}

// Subdomains related to one or more domains (using public APIs)
func Subdomainer(ctx *fasthttp.RequestCtx, client *fasthttp.Client, threads int) {
	// incoming request
	var params subdomainerParams
	if json.Unmarshal(ctx.PostBody(), &params) != nil {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// provider
	provider, ok := subdomainerProviders[params.Provider]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// content type
	ctx.SetContentType("application/json; charset=utf-8")
	// concurrency
	var warns []string = []string{}
	var subdomains []string = []string{}
	var lock sync.Mutex = sync.Mutex{}
	var pool chan struct{} = make(chan struct{}, threads)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(params.Domains))
	// iteration
	for _, domain := range params.Domains {
		pool <- struct{}{}
		go func(domain string) {
			// action
			_, body, err := client.Get(nil, fmt.Sprintf(provider.Endpoint(client), domain))
			if err != nil {
				lock.Lock()
				warns = append(warns, err.Error())
				lock.Unlock()
			} else {
				// response
				lock.Lock()
				subdomains = append(subdomains, provider.Get(body)...)
				lock.Unlock()
			}
			<-pool
			wg.Done()
		}(domain)
	}
	wg.Wait()
	close(pool)
	// outgoing response
	var response subdomainerResponse = subdomainerResponse{
		Domains:    params.Domains,
		Subdomains: helpers.UniqStrings(subdomains),
		Warnings:   warns,
	}
	result, _ := json.Marshal(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "%s", result)
}
