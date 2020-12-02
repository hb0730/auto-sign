package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	Method string
	Url    string
	Params string
}
type Cookies map[string]string

type AutoSign interface {
	Do()
}

func (rq *Request) CreateRequest() *http.Request {
	if rq.Url == "" {
		fmt.Println("request url is null")
		return nil
	}
	if rq.Method == "" {
		rq.Method = "GET"
	}
	request, e := http.NewRequest(rq.Method, rq.Url, strings.NewReader(rq.Params))
	if e != nil {
		fmt.Printf("http.NewRequest %v\n", e)
		return nil
	}
	return request
}

func SetCookie(cookies map[string]string, request *http.Request) {
	for k, v := range cookies {
		cookie := http.Cookie{Name: k, Value: v}
		request.AddCookie(&cookie)
	}
}
func Query(method string, url string, params string, cookies Cookies) (string, bool) {
	r := Request{Method: method, Url: url, Params: params}
	rq := r.CreateRequest()
	return Req(rq, cookies)
}
func Req(request *http.Request, cookies Cookies) (string, bool) {
	if request == nil {
		fmt.Println("request failed")
		return "", false
	}
	SetCookie(cookies, request)
	body, _, is := ClientDo(request)
	if is {
		return body, true

	}
	return "", false
}

func ClientDo(request *http.Request) (string, []*http.Cookie, bool) {
	response, e := http.DefaultClient.Do(request)
	if e != nil {
		fmt.Printf("request error %v\n", e)
		return "", nil, false
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), response.Cookies(), true
}
