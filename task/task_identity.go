package task

import (
	"encoding/json"
	"errors"

	"github.com/MasterJoyHunan/flowablesdk"
	"github.com/MasterJoyHunan/flowablesdk/candidate"
	"github.com/MasterJoyHunan/flowablesdk/pkg/httpclient"
)

// ListIdentity 获取任务的所有相关人员(包括组和人)
func (t Tasks) ListIdentity(taskId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListIdentityApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// ListUsersIdentity 获取任务的所有相关用户
func (t Tasks) ListUsersIdentity(taskId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListUsersIdentityApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// ListGroupsIdentity 获取任务的所有相关组
func (t Tasks) ListGroupsIdentity(taskId string) (resp []candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(ListGroupsIdentityApi, taskId)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DetailIdentity 获取任务特定相关人员
func (t Tasks) DetailIdentity(taskId string, req DetailIdentityRequest) (resp candidate.Candidate, err error) {
	if !(req.Family == "groups" || req.Family == "users") {
		err = errors.New("family not in groups or users")
	}
	request := flowablesdk.GetRequest(DetailIdentityApi, taskId, req.Family, req.IdentityId, req.Type)
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// AddIdentity 任务添加相关人员
func (t Tasks) AddIdentity(taskId string, req AddIdentityRequest) (resp candidate.Candidate, err error) {
	request := flowablesdk.GetRequest(AddIdentityApi, taskId)
	request.With(httpclient.WithJson(req))
	data, err := request.DoHttpRequest()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &resp)
	return
}

// DeleteIdentity 任务添加相关人员
func (t Tasks) DeleteIdentity(taskId string, req DetailIdentityRequest) error {
	if !(req.Family == "groups" || req.Family == "users") {
		return errors.New("family not in groups or users")
	}
	request := flowablesdk.GetRequest(DeleteIdentityApi, taskId, req.Family, req.IdentityId, req.Type)
	_, err := request.DoHttpRequest()
	return err
}
