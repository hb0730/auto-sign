package request

import (
	"compress/gzip"
	"fmt"
	"github.com/hb0730/auto-sign/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	request  *http.Request
	client   *http.Client
	response *http.Response
}

func CreateRequest(method, url, params string) (*Request, error) {
	re, err := http.NewRequest(method, url, strings.NewReader(params))
	if err != nil {
		return nil, utils.AutoSignError{
			Module:  "request",
			Method:  "create Request",
			Message: fmt.Sprintf("createRequest error ,error messge: [%s]", err.Error()),
		}
	}
	request := new(Request)
	request.request = re
	return request, nil
}
func (r *Request) Header(header http.Header) {
	r.request.Header = header
}

func (r *Request) SetHeader(k, v string) {
	r.request.Header.Set(k, v)
}
func (r *Request) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		r.SetHeader(k, v)
	}
}
func (r *Request) AddHeader(k, v string) {
	r.request.Header.Add(k, v)
}

func (r *Request) AddHeaders(headers map[string]string) {
	for k, v := range headers {
		r.AddHeader(k, v)
	}
}

func (r *Request) AddCookie(cookie *http.Cookie) {
	r.request.AddCookie(cookie)
}

func (r *Request) AddCookies(cookies []*http.Cookie) {
	for _, k := range cookies {
		r.AddCookie(k)
	}
}

func (r *Request) AddCookieFromNameValue(name, value string) {
	c := &http.Cookie{
		Name:  name,
		Value: value,
	}
	r.AddCookie(c)
}

func (r *Request) AddCookiesFromMap(c map[string]string) {
	cookies := make([]*http.Cookie, 0)
	for k, v := range c {
		cookie := &http.Cookie{
			Name:  k,
			Value: v,
		}
		cookies = append(cookies, cookie)
	}
	r.AddCookies(cookies)
}

func (r *Request) GetRequest() *http.Request {
	return r.request
}
func (r *Request) SetClient(client *http.Client) {
	r.client = client
}

func (r *Request) Do() error {
	if r.client == nil {
		r.client = http.DefaultClient
	}
	response, err := r.client.Do(r.request)
	if err != nil {
		return utils.AutoSignError{
			Module:  "request",
			Method:  "Client DO",
			Message: fmt.Sprintf("ClientDo Error,Error message: [%s]", err.Error()),
		}
	}
	r.response = response
	return nil
}

func (r *Request) GetBody() ([]byte, error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.response.Body)
	var reader io.Reader
	var err error
	if r.response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(r.response.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = r.response.Body
	}
	return ioutil.ReadAll(reader)
}

// ConvertHeader 将 map[string]string 转成 http.header
//   headers 可以为空
func ConvertHeader(header http.Header, headers map[string]string) http.Header {
	if header == nil {
		header = http.Header{}
	}
	for k, v := range headers {
		header[k] = []string{v}
	}
	return header
}
