package spiders

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// CommonCrawlEndpoint returns the CommonCrawl API path
func CommonCrawlEndpoint(client *fasthttp.Client, includeSubdomains bool) string {
	var endpoint string
	if includeSubdomains {
		endpoint = "https://index.commoncrawl.org/CC-MAIN-2021-21-index?format=json&url=*.%s"
	} else {
		endpoint = "https://index.commoncrawl.org/CC-MAIN-2021-21-index?format=json&url=%s"
	}
	_, body, err := client.Get(nil, "http://index.commoncrawl.org/collinfo.json")
	if err != nil {
		return endpoint
	}
	var api []struct {
		Endpoint string `json:"cdx-api"`
	}
	if err = json.Unmarshal(body, &api); err != nil {
		return endpoint
	}
	if len(api) == 0 {
		return endpoint
	} else if includeSubdomains {
		return fmt.Sprintf("%s?format=json&url=*.%s", api[0].Endpoint, "%s")
	} else {
		return fmt.Sprintf("%s?format=json&url=%s", api[0].Endpoint, "%s")
	}
}

// CommonCrawlResult returns the query performed against the CommonCrawl API
func CommonCrawlResult(body []byte) []string {
	var links []string = []string{}
	var wrapper []struct {
		URL string `json:"url"`
	}
	if json.Unmarshal(body, &wrapper) == nil {
		links = make([]string, 0, len(wrapper))
		for i := 0; i < len(wrapper); i++ {
			links = append(links, wrapper[i].URL)
		}
	}
	return links
}
