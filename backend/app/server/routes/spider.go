package routes

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/valyala/fasthttp"

	"cosasdepuma/masterchef/app/helpers"
	"cosasdepuma/masterchef/app/providers/spiders"
)

type (
	spiderProvider struct {
		Endpoint func(*fasthttp.Client, bool) string
		Get      func([]byte) []string
	}
	spiderParams struct {
		Domains           []string `json:"domains"`
		Provider          string   `json:"provider"`
		IncludeSubdomains bool     `json:"includeSubdomains"`
	}

	spiderResponse struct {
		Domains  []string `json:"domains"`
		Links    []string `json:"links"`
		Warnings []string `json:"warnings"`
	}
)

var spiderProviders map[string]spiderProvider = map[string]spiderProvider{
	"alienvault": {
		Endpoint: spiders.AlienVaultEndpoint,
		Get:      spiders.AlienVaultResult,
	},
	"commoncrawl": {
		Endpoint: spiders.CommonCrawlEndpoint,
		Get:      spiders.CommonCrawlResult,
	},
	"waybackmachine": {
		Endpoint: spiders.WaybackMachineEndpoint,
		Get:      spiders.WaybackMachineResult,
	},
}

// Spider returns links related to one or more domains (using public APIs)
func Spider(ctx *fasthttp.RequestCtx, client *fasthttp.Client, threads int) {
	// incoming request
	var params spiderParams
	if json.Unmarshal(ctx.PostBody(), &params) != nil {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// provider
	provider, ok := spiderProviders[params.Provider]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// content type
	ctx.SetContentType("application/json; charset=utf-8")
	// concurrency
	var warns []string = []string{}
	var links []string = []string{}
	var lock sync.Mutex = sync.Mutex{}
	var pool chan struct{} = make(chan struct{}, threads)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(params.Domains))
	// iteration
	for _, domain := range params.Domains {
		pool <- struct{}{}
		go func(domain string) {
			// action
			_, body, err := client.Get(nil, fmt.Sprintf(provider.Endpoint(client, params.IncludeSubdomains), domain))
			if err != nil {
				lock.Lock()
				warns = append(warns, err.Error())
				lock.Unlock()
			} else {
				// response
				lock.Lock()
				links = append(links, provider.Get(body)...)
				lock.Unlock()
			}
			<-pool
			wg.Done()
		}(domain)
	}
	wg.Wait()
	close(pool)
	// outgoing response
	var response spiderResponse = spiderResponse{
		Domains:  params.Domains,
		Links:    helpers.UniqStrings(links),
		Warnings: warns,
	}
	result, _ := json.Marshal(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "%s", result)
}
