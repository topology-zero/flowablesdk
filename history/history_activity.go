package history

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// ListActivity 流程实例历史流转
func (h History) ListActivity(req ListActivityRequest) (resp ListActivityResponse, err error) {
	request := flowablesdk.GetRequest(ListActivityApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
