package history

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// ListVariable 流程实例历史流转
func (h History) ListVariable(req ListVariableRequest) (resp ListVariableResponse, err error) {
	request := flowablesdk.GetRequest(ListVariableApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
