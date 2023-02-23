package task_events

import (
	"encoding/json"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/events"
)

// List 获取任务的所有操作事件
func List(taskId string) (resp []events.Events, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取任务特定操作事件
func Detail(taskId, eventsId string) (resp events.Events, err error) {
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
