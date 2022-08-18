package flowablesdk

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

func GetRequest(api *Api, params ...any) *httpclient.Request {
	url := Configs.Url + api.Url
	if len(params) > 0 {
		url = fmt.Sprintf(url, params...)
	}
	return httpclient.NewHttpRequest(
		api.Method,
		url,
		httpclient.WithTimeout(15*time.Second),
		httpclient.WithHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(Configs.Username+":"+Configs.Password))),
	)
}
