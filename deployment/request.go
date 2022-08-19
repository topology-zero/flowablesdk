package deployment

type ListRequest struct {
	Name     string
	Category string
	Sort     string
}

type CreateRequest struct {
	FileName string // must end by .bpmn20.xml
	Xml      string
}
