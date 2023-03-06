package task_event

import (
	"encoding/json"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/event"
)

// List 获取任务的所有操作事件
func List(taskId string) (resp []event.Event, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取任务特定操作事件
func Detail(taskId, eventsId string) (resp event.Event, err error) {
	request := flowablesdk.GetRequest(DetailApi, taskId, eventsId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 获取任务特定操作事件 (谨慎使用)
func Delete(taskId, eventsId string) (err error) {
	request := flowablesdk.GetRequest(DeleteEventsApi, taskId, eventsId)
	_, err = request.DoHttpRequest()
	return
}
