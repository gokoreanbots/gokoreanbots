package gokoreanbots

import (
	"fmt"
	"net/url"
	"strings"
)

func urlFmt(uri string, route string, params *strMap) string {
	var paramList []string
	if params != nil {
		for k, v := range *params {
			paramList = append(paramList, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
	}
	return fmt.Sprintf("%s%s?%s", uri, route, strings.Join(paramList, "&"))
}
