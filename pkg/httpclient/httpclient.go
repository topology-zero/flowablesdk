package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Method string

const (
	GET    = Method("GET")
	POST   = Method("POST")
	PUT    = Method("PUT")
	DELETE = Method("DELETE")

	FlowableUserId = "FlowableUserId"
)

type Request struct {
	method string
	url    string
	body   io.Reader
	header map[string]string
	query  map[string]string

	requestDebug  bool
	responseDebug bool

	client *http.Client
}

type HTTPOption func(r *Request)

// NewHttpRequest 构建一个 Request 对象
func NewHttpRequest(m Method, url string, opt ...HTTPOption) *Request {
	r := &Request{
		method: string(m),
		header: make(map[string]string),
		query:  make(map[string]string),
		url:    url,
		client: &http.Client{},
	}
	return r.With(opt...)
}

func (r *Request) With(opt ...HTTPOption) *Request {
	for _, option := range opt {
		option(r)
	}
	return r
}

// DoHttpRequest 发送HTTP请求
func (r *Request) DoHttpRequest() ([]byte, error) {
	request, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.header {
		request.Header.Add(k, v)
	}

	if r.query != nil && len(r.query) > 0 {
		for k, v := range r.query {
			query := request.URL.Query()
			query.Add(k, v)
			request.URL.RawQuery = query.Encode()
		}
	}

	r.RequestDebug(request)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}

	respStr, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	r.ResponseDebug(response.StatusCode, respStr)

	if response.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("status code = %d , response data = %s", response.StatusCode, string(respStr)))
	}

	return respStr, nil
}

func (r *Request) RequestDebug(request *http.Request) {
	if !r.requestDebug {
		return
	}

	log.Println("[", r.method, "]", r.url)
	if r.method == "GET" {
		log.Println("query = ", request.URL.RawQuery)
	} else {
		ct := request.Header.Get("Content-Type")
		if strings.HasPrefix(ct, "application/json") {
			body, _ := ioutil.ReadAll(request.Body)
			request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			log.Println("request = ", string(body))
		} else if strings.HasPrefix(ct, "application/x-www-form-urlencoded") {
			request.ParseForm()
			log.Println("request = ", request.PostForm.Encode())
		} else if strings.HasPrefix(ct, "multipart/form-data") {
			request.ParseMultipartForm(10e6)
			log.Println("request = ", request.PostForm.Encode())
		} else {
			log.Println("not debug , Content-Type = ", ct)
		}
	}
}

func (r *Request) ResponseDebug(code int, resp []byte) {
	if !r.responseDebug {
		return
	}
	log.Printf("code = [%d] reps = %s\n", code, string(resp))
}

func WithTimeout(d time.Duration) HTTPOption {
	return func(r *Request) {
		r.client.Timeout = d
	}
}

func WithQuery(query map[string]string) HTTPOption {
	return func(r *Request) {
		r.query = query
	}
}

func WithJson(body any) HTTPOption {
	return func(r *Request) {
		marshal, _ := json.Marshal(body)
		r.body = bytes.NewBuffer(marshal)
		r.header["Content-Type"] = "application/json;charset=UTF-8"
	}
}

func WithFromData(body map[string]any) HTTPOption {
	return func(r *Request) {
		r.body = strings.NewReader(Map2Str(body))
		r.header["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8"
	}
}

func WithRequestDebug() HTTPOption {
	return func(r *Request) {
		r.requestDebug = true
	}
}

func WithResponseDebug() HTTPOption {
	return func(r *Request) {
		r.responseDebug = true
	}
}

func WithMultipartFrom(file *UploadFile, fromData map[string]string) HTTPOption {
	body := &bytes.Buffer{}
	newWriter := multipart.NewWriter(body)

	// 文件处理
	if file != nil {
		formFile, _ := newWriter.CreateFormFile(file.Field, file.FileName)
		_, _ = io.Copy(formFile, file.File)
	}

	// 普通字段处理
	if fromData != nil && len(fromData) > 0 {
		for k, v := range fromData {
			_ = newWriter.WriteField(k, v)
		}
	}

	_ = newWriter.Close()

	return func(r *Request) {
		r.body = body
		r.header["Content-Type"] = newWriter.FormDataContentType()
	}
}

func WithHeader(key, val string) HTTPOption {
	return func(r *Request) {
		r.header[key] = val
	}
}

func Map2Str(m map[string]any) string {
	var strArr []string
	for k, v := range m {
		strArr = append(strArr, k+"="+ConvertString(v))
	}
	return strings.Join(strArr, "&")
}

func ConvertString(value any) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case int64:
		return strconv.Itoa(int(value))
	case json.Number:
		return value.String()
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	default:
		return ""
	}
}
