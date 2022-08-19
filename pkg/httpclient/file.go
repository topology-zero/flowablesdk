package httpclient

import "io"

type UploadFile struct {
	Field    string    // fromData 字段名
	FileName string    // 文件名
	File     io.Reader // 文件
}
