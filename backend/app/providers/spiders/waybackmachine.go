package spiders

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// WaybackMachineEndpoint returns the WayBackMachine API path
func WaybackMachineEndpoint(_ *fasthttp.Client, includeSubdomains bool) string {
	if includeSubdomains {
		return "http://web.archive.org/cdx/search/cdx?output=json&collapse=urlkey&url=*.%s/*"
	} else {
		return "http://web.archive.org/cdx/search/cdx?output=json&collapse=urlkey&url=%s/*"
	}
}

// WaybackMachineResult returns the query performed against the WayBackMachine API
func WaybackMachineResult(body []byte) []string {
	var links []string = []string{}
	var wrapper [][]string
	if json.Unmarshal(body, &wrapper) == nil {
		links = make([]string, 0, len(wrapper))
		for i := 1; i < len(wrapper); i++ {
			if wrapper[i][4] != "404" {
				links = append(links, wrapper[i][2])
			}
		}
	}
	return links
}
