package variable

type Variable struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    any    `json:"value"`
	Scope    string `json:"scope"`
	ValueUrl string `json:"valueUrl"`
}
