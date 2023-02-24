package model_source

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/flowablesdk"
)

func TestMain(m *testing.M) {
	flowablesdk.Setup(flowablesdk.Config{Url: "http://127.0.0.1:8080/"})
	m.Run()
}

const id = "b785a233-b413-11ed-b7d1-38f3ab6b92c1"

//go:embed _js3.png
var image []byte

func TestDetailSource(t *testing.T) {
	data, err := DetailSource(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(data))
}

func TestUpdateSource(t *testing.T) {
	err := UpdateSource(id, SourceRequest{
		FileName: "UpdateSource.png",
		Value:    bytes.NewReader(image),
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("upload success")
}

func TestDetailExtraSource(t *testing.T) {
	data, err := DetailExtraSource(id)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(data))
}

func TestUpdateExtraSource(t *testing.T) {
	err := UpdateExtraSource(id, SourceRequest{
		FileName: "UpdateExtraSource.png",
		Value:    bytes.NewReader(image),
	})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("upload success")
}
