package execution

import (
	"encoding/json"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/common"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

type Execution struct {
	Id                 string `json:"id"`
	Url                string `json:"url"`
	ParentId           string `json:"parentId"`
	ParentUrl          string `json:"parentUrl"`
	ProcessInstanceId  string `json:"processInstanceId"`
	ProcessInstanceUrl string `json:"processInstanceUrl"`
	Suspended          bool   `json:"suspended"`
	ActivityId         string `json:"activityId"`
	TenantId           string `json:"tenantId"`
}

// List 获取流程执行
func List(req ListRequest) (list []Execution, count int, err error) {
	request := flowablesdk.GetRequest(ListApi)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	var commonData common.ListCommonResponse
	err = json.Unmarshal(data, &commonData)
	if err != nil {
		return
	}

	count = commonData.Total
	err = json.Unmarshal(commonData.Data, &list)
	return
}

// Detail 获取单个流程执行
func Detail(executionId string) (resp Execution, err error) {
	request := flowablesdk.GetRequest(DetailApi, executionId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// ListActive Get active activities in an execution
// Returns all activities which are active in the execution and
// in all child-executions (and their children, recursively), if any.
func ListActive(executionId string) (resp []string, err error) {
	request := flowablesdk.GetRequest(ActiveApi, executionId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Execute execute an action on an execution
func Execute(executionId string, req ExecuteRequest) (resp []byte, err error) {
	request := flowablesdk.GetRequest(ExecuteApi, executionId)
	request.With(httpclient.WithJson(req))
	// TODO 返回值有可能是 ExecuteResponse, 也有可能是 Execution
	return request.DoHttpRequest()
}
