package flowablesdk

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

func GetRequest(api *Api, params ...any) *httpclient.Request {
	url := Configs.Url
	if api.Prefix == ProcessPrefix {
		url += Configs.ProcessPrefix
	} else if api.Prefix == FormPrefix {
		url += Configs.FormPrefix
	}
	url += api.Url

	if len(params) > 0 {
		url = fmt.Sprintf(url, params...)
	}
	request := httpclient.NewHttpRequest(
		api.Method,
		url,
		httpclient.WithTimeout(15*time.Second),
		httpclient.WithHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(Configs.Username+":"+Configs.Password))),
	)

	if Configs.RequestDebug {
		request.With(httpclient.WithRequestDebug())
	}

	if Configs.ResponseDebug {
		request.With(httpclient.WithResponseDebug())
	}

	return request
}
