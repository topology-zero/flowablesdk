package deployment

type ListRequest struct {
	Name     string
	Category string
	Sort     string
}

type CreateRequest struct {
	FileName string
	Xml      string
}
