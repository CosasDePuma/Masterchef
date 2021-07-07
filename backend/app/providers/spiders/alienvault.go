package spiders

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// AlienVaultEndpoint returns the AlienVault API path
func AlienVaultEndpoint(_ *fasthttp.Client, includeSubdomains bool) string {
	if includeSubdomains {
		return "https://otx.alienvault.com/api/v1/indicators/domain/%s/url_list?limit=50"
	} else {
		return "https://otx.alienvault.com/api/v1/indicators/hostname/%s/url_list?limit=50"
	}
}

// AlienVaultResult returns the query performed against the AlienVault API
func AlienVaultResult(body []byte) []string {
	var links []string = []string{}
	var wrapper struct {
		URLs []struct {
			URL string `json:"url"`
		} `json:"url_list"`
	}
	if json.Unmarshal(body, &wrapper) == nil {
		links = make([]string, 0, len(wrapper.URLs))
		for i := 0; i < len(wrapper.URLs); i++ {
			links = append(links, wrapper.URLs[i].URL)
		}
	}
	return links
}
