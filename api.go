package flowablesdk

import "github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"

type Api struct {
	Method httpclient.Method
	Url    string
}

func NewApi(m httpclient.Method, url string) *Api {
	return &Api{
		Method: m,
		Url:    url,
	}
}
