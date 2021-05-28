package utils

import (
	"compress/gzip"
	"github.com/mritd/logger"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//Cookies request cookie
type Cookies map[string]string

//Request utils
type Request struct {
	// Method 请求方法
	// Get/Post
	Method string
	// Url 请求路径
	Url string
	// Params 请求参数
	Params string
}

// CreateRequest  创建请求
// 返回http.Request与error
func (r Request) CreateRequest() (*http.Request, error) {
	if r.Url == "" {
		logger.Warn("[request] Request Url must be not null")
		return nil, &AutoSignError{Module: "request utils", Method: "create request error", Message: "create request error :Request Url must be not null"}
	}
	if r.Method == "" {
		r.Method = "GET"
	}
	request, err := http.NewRequest(r.Method, r.Url, strings.NewReader(r.Params))
	if err != nil {
		logger.Errorf("[request] http.NewRequest %v\n", err)
		return nil, &AutoSignError{
			Method:  "requst utils",
			Message: "NewRequest error",
			E:       err,
		}
	}
	return request, nil
}

// HttpRequest request请求并设置Cookies
// 如果不存在Cookie可以直接使用ClientDo
func HttpRequest(request *http.Request, cookies Cookies) (*http.Response, error) {
	cs := ConvertCookies(cookies)
	for _, v := range cs {
		request.AddCookie(v)
	}
	return ClientDo(request)
}

// ClientDo 执行request请求获取response响应
func ClientDo(request *http.Request) (*http.Response, error) {
	response, e := http.DefaultClient.Do(request)
	if e != nil {
		logger.Error("[request] request error %v\n", e)
		return nil, &AutoSignError{
			Module: "request utils",
			Method: "ClientDo error",
			E:      e,
		}
	}
	return response, nil
}

// ConvertCookies 将Cookies转换成http.Cookie
func ConvertCookies(cookies Cookies) []*http.Cookie {
	var array = make([]*http.Cookie, 0)
	for k, v := range cookies {
		a := &http.Cookie{Name: k, Value: v}
		array = append(array, a)
	}
	return array
}

// GetBody 读取Response内容
func GetBody(response *http.Response) ([]byte, error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	var reader io.Reader
	var err error
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = response.Body
	}
	return ioutil.ReadAll(reader)
}
