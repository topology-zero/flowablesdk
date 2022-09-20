package history

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHistory_ListComment(t *testing.T) {
	var his History
	data, err := his.ListComment(id)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestHistory_AddComment(t *testing.T) {
	var his History
	data, err := his.AddComment(id, AddCommentRequest{
		Message:               "hello world",
		SaveProcessInstanceId: false,
	})
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
