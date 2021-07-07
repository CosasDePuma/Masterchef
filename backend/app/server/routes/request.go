package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"

	"cosasdepuma/masterchef/app/server/client"
)

type (
	requestParams = struct {
		Method  string            `json:"method"`
		URLs    []string          `json:"urls"`
		Headers map[string]string `json:"headers"`
		Proxy   string            `json:"proxy"`
	}
	requestResponses = struct {
		URLs      []string          `json:"urls"`
		Responses []requestResponse `json:"responses"`
		Warnings  []string          `json:"warnings"`
	}
	requestResponse = struct {
		URL     string            `json:"url"`
		Status  int               `json:"status"`
		Headers map[string]string `json:"headers"`
		Length  int               `json:"length"`
		Body    []byte            `json:"body"`
	}
)

// Request to one or more websites
func Request(ctx *fasthttp.RequestCtx, clnt *fasthttp.Client, threads int) {
	// incoming request
	var params requestParams
	if json.Unmarshal(ctx.PostBody(), &params) != nil {
		ctx.SetStatusCode(fasthttp.StatusNotAcceptable)
		return
	}
	// proxy
	if params.Proxy != "" {
		proxy := strings.SplitN(params.Proxy, "://", 2)
		if len(proxy) == 2 {
			clnt = client.New()
			switch proxy[0] {
			case "http":
				clnt.Dial = fasthttpproxy.FasthttpHTTPDialer(proxy[1])
			case "socks4", "socks5":
				clnt.Dial = fasthttpproxy.FasthttpSocksDialer(proxy[1])
			}
		}
	}
	// content type
	ctx.SetContentType("application/json; charset=utf-8")
	// concurrency
	var warns []string = []string{}
	var responses []requestResponse = []requestResponse{}
	var lock sync.Mutex = sync.Mutex{}
	var pool chan struct{} = make(chan struct{}, threads)
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(len(params.URLs))
	// iteration
	for _, url := range params.URLs {
		pool <- struct{}{}
		go func(url string) {
			// outgoing requests
			var req *fasthttp.Request = fasthttp.AcquireRequest()
			defer fasthttp.ReleaseRequest(req)
			req.SetRequestURI(url)
			req.Header.SetMethod(params.Method)
			req.Header.Set("Connection", "close")
			req.Header.Set("Accept-Encoding", "gzip")
			for key, val := range params.Headers {
				req.Header.Set(key, val)
			}
			// response
			var res *fasthttp.Response = fasthttp.AcquireResponse()
			defer fasthttp.ReleaseResponse(res)
			// action
			if err := clnt.DoRedirects(req, res, 10); err != nil {
				lock.Lock()
				warns = append(warns, err.Error())
				lock.Unlock()
			} else {
				// body
				var body []byte
				if bytes.EqualFold(res.Header.Peek("Content-Encoding"), []byte("gzip")) {
					body, _ = res.BodyGunzip()
				} else {
					body = res.Body()
				}
				// headers
				var headers map[string]string = map[string]string{}
				res.Header.VisitAll(func(k []byte, v []byte) {
					headers[string(k)] = string(v)
				})
				// result
				responses = append(responses, requestResponse{
					URL:     url,
					Status:  res.StatusCode(),
					Length:  res.Header.ContentLength(),
					Headers: headers,
					Body:    body,
				})
			}

			<-pool
			wg.Done()
		}(url)
	}
	wg.Wait()
	close(pool)
	// outgoing response
	var response requestResponses = requestResponses{
		URLs:      params.URLs,
		Responses: responses,
		Warnings:  warns,
	}
	result, _ := json.Marshal(response)
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "%s", result)
}
