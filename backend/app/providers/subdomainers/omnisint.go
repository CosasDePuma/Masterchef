package subdomainers

import (
	"bytes"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// OmnisintEndpoint returns the Omnisint API path
func OmnisintEndpoint(_ *fasthttp.Client) string {
	return "https://sonar.omnisint.io/subdomains/%s"
}

// OmnisintResult returns the query performed against the Omnisint API
func OmnisintResult(body []byte) []string {
	var subdomains []string = []string{}
	if !bytes.EqualFold(body, []byte("null\n")) {
		_ = json.Unmarshal(body, &subdomains)
	}
	return subdomains
}
