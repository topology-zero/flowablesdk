package history

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHistory_ListVariable(t *testing.T) {
	var his History
	req := ListVariableRequest{}
	req.ProcessInstanceId = id
	data, err := his.ListVariable(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
