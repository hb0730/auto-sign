package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Method  string
	Url     string
	Params  string
	Headers http.Header
}
type AutoSign interface {
	Do()
}

func (rq *Request) Request() (string, []*http.Cookie, bool) {
	if rq.Url == "" {
		fmt.Println("request url is null")
		return "", nil, false
	}
	if rq.Method == "" {
		rq.Method = "GET"
	}

	request, e := http.NewRequest(rq.Method, rq.Url, strings.NewReader(rq.Params))
	if e != nil {
		fmt.Printf("http.NewRequest %v\n", e)
		return "", nil, false
	}
	request.Header = rq.Headers
	response, e := http.DefaultClient.Do(request)
	if e != nil {
		fmt.Printf("request error %v\n", e)
		return "", nil, false
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), response.Cookies(), true
}
