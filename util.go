package gapi

import (
	"fmt"
	"net/url"
)

func buildPathAndQuery(path string, data map[string]string) string {
	pathAndQuery := fmt.Sprintf("%s?", path)
	params := url.Values{}

	for k, v := range data {
		params.Add(k, v)
	}

	return pathAndQuery + params.Encode()
}
