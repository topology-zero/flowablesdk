package execution

import "github.com/MasterJoyHunan/flowablesdk/variable"

type ExecuteResponse struct {
	Action     string                    `json:"action"`
	SignalName string                    `json:"signalName"`
	Variables  []variable.SimpleVariable `json:"variables"`
}
