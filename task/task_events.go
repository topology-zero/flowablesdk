package task

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/events"
)

// ListEvents 获取任务的所有相关组
func (t Tasks) ListEvents(taskId string) (resp []events.Events, err error) {
	request := flowablesdk.GetRequest(ListEventsApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DetailEvents 获取任务特定相关人员
func (t Tasks) DetailEvents(taskId, eventsId string) (resp events.Events, err error) {
	request := flowablesdk.GetRequest(DetailEventsApi, taskId, eventsId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}
