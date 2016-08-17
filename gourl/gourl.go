package main

import (
	"github.com/astaxie/beego/plugins/apiauth"
	"net/url"
	"fmt"
)

func main() {
	//appsecret, method string, params url.Values, RequestURI string
	appsecret := "1"
	method := "GET"
	url := url.URL{
		Scheme:   "http",
		Host:     "http://127.0.0.1:8080/",
		Path:     "/api",
		RawQuery: "appid=1",
		Fragment: "#",
	}
	fmt.Println(url.String())
	fmt.Println(url.RequestURI())
	fmt.Println(apiauth.Signature(appsecret, method, url.Query(),url.RequestURI()))
}
