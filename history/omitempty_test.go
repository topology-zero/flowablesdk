package history

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Omitempty struct {
	Time  time.Time         `json:"time,omitempty"`  // 不是指针方式,无法忽略
	Slice []string          `json:"slice,omitempty"` // 空值可以忽略
	B     bool              `json:"b,omitempty"`     // 如果是 false, 可以忽略
	S     []S               `json:"s,omitempty"`     // 空值可以忽略
	M     map[string]string `json:"m,omitempty"`     // 空值可以忽略
}

type S struct {
	S1 string
	M2 map[string]string
}

func TestOmitempty(t *testing.T) {
	data := Omitempty{}
	jsonStr, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println(string(jsonStr))
}
