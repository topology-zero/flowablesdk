package history

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHistory_ListActivity(t *testing.T) {
	var his History
	req := ListActivityRequest{
		//ProcessInstanceId: id,
	}
	req.Size = 3
	data, err := his.ListActivity(req)
	if err != nil {
		t.Error(err)
		return
	}

	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
