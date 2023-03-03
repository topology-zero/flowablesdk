package model

type Model struct {
	Name                string        `json:"name,omitempty"`
	Key                 string        `json:"key,omitempty"`
	Version             string        `json:"version,omitempty"`
	Description         string        `json:"description,omitempty"`
	Url                 string        `json:"url,omitempty"`
	Fields              []FormField   `json:"fields,omitempty"`
	Outcomes            []FormOutcome `json:"outcomes,omitempty"`
	OutcomeVariableName string        `json:"outcomeVariableName,omitempty"`
}

type FormField struct {
	Id          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Type        string         `json:"type,omitempty"`
	Value       any            `json:"value,omitempty"`
	Required    bool           `json:"required,omitempty"`
	ReadOnly    bool           `json:"readOnly,omitempty"`
	OverrideId  bool           `json:"overrideId,omitempty"`
	Placeholder string         `json:"placeholder,omitempty"`
	Params      map[string]any `json:"params,omitempty"`
}

type FormOutcome struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
