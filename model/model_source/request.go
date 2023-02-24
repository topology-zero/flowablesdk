package model_source

import "io"

type SourceRequest struct {
	FileName string
	Value    io.Reader
}
