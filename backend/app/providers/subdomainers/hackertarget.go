package subdomainers

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/valyala/fasthttp"
)

// HackerTargetEndpoint returns the HackerTarget API path
func HackerTargetEndpoint(_ *fasthttp.Client) string {
	return "https://api.hackertarget.com/hostsearch/?q=%s"
}

// HackerTargetResult returns the query performed against the HackerTarget API
func HackerTargetResult(body []byte) []string {
	var subdomains []string = []string{}
	var stream *bufio.Scanner = bufio.NewScanner(bytes.NewBuffer(body))
	for stream.Scan() {
		var data []string = strings.SplitN(stream.Text(), ",", 2)
		if len(data) == 2 {
			subdomains = append(subdomains, data[0])
		}
	}
	return subdomains
}
