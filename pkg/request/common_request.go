package request

type CommonRequest struct {
	Start int    // 起始页
	Size  int    // 每页显示多少条
	Sort  string // 排序字段
	Order string // 排序方法 asc | desc
}
