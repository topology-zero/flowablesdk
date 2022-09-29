package model

type Model struct {
	Url                 string
	Fields              []FormField
	Outcomes            []FormOutcome
	OutcomeVariableName string
}

type FormField struct {
	Id          string
	Name        string
	Type        string
	Value       any
	Required    bool
	ReadOnly    bool
	OverrideId  bool
	Placeholder string
	Params      map[string]any
}

type FormOutcome struct {
	Id   string
	Name string
}
