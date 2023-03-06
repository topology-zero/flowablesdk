package task_identity

import (
	"encoding/json"
	"errors"

	"github.com/topology-zero/flowablesdk"
	"github.com/topology-zero/flowablesdk/candidate"
	"github.com/topology-zero/flowablesdk/pkg/httpclient"
)

// List 获取任务的所有相关人员(包括组和人)
func List(taskId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// ListUsers 获取任务的所有相关用户
func ListUsers(taskId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListUsersApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// ListGroups 获取任务的所有相关组
func ListGroups(taskId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListGroupsApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Detail 获取任务特定相关人员
func Detail(taskId string, req DetailRequest) (resp candidate.Candidate, err error) {
	if !(req.Family == "groups" || req.Family == "users") {
		err = errors.New("family not in groups or users")
	}
	request := flowablesdk.GetRequest(DetailApi, taskId, req.Family, req.IdentityId, req.Type)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Add 任务添加相关人员
//
// 添加人到任务, type 可以自定义
//{
//  "userId" : "kermit",
//  "type" : "candidate",
//}
//
// 添加组到任务, type 可以自定义
//{
//  "groupId" : "sales",
//  "type" : "candidate",
//}
func Add(taskId string, req AddRequest) (resp candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(AddApi, taskId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// Delete 任务删除相关人员
func Delete(taskId string, req DetailRequest) error {
	if !(req.Family == "groups" || req.Family == "users") {
		return errors.New("family not in groups or users")
	}
	request := flowablesdk.GetRequest(DeleteApi, taskId, req.Family, req.IdentityId, req.Type)
	_, err := request.DoHttpRequest()
	return err
}
