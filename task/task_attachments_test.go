package task

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

//go:embed js3.png
var file string

const attachmentsId = "526cc4e9-21f4-11ed-b5b7-0242ac140002"

func TestTasks_ListAttachments(t *testing.T) {
	var task Tasks
	data, err := task.ListAttachments(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_AddAttachmentsUrl(t *testing.T) {
	var task Tasks
	data, err := task.AddAttachmentsUrl(id, AddAttachmentsUrlRequest{
		Name:        "Simple attachment",
		Description: "Simple attachment description",
		Type:        "simpleType",
		ExternalUrl: "http://www.baidu.org",
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_AddAttachments(t *testing.T) {
	var task Tasks
	data, err := task.AddAttachments(id, AddAttachmentsRequest{
		Name:        "Simple attachment1",
		Description: "Simple attachment description1",
		Type:        "simpleType1",
		File:        strings.NewReader(file),
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestTasks_DeleteAttachments(t *testing.T) {
	var task Tasks
	err := task.DeleteAttachments(id, attachmentsId, "1")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("delete success")
}

func TestTasks_ContentAttachments(t *testing.T) {
	var task Tasks
	data, err := task.ContentAttachments(id, "072a57ad-3883-11ed-9f1e-38f3ab6b92c1")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(data)
}
