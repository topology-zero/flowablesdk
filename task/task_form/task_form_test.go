package task_form

import (
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const id = "2a0c185b-b80b-11ed-ac5d-38f3ab6b92c1"

func TestGetForm(t *testing.T) {
	str, err := GetForm(id)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(str))
}
