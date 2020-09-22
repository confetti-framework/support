package support

import (
	"bytes"
	"mime/multipart"
)

type Files []File

type File struct {
	source multipart.File
	header *multipart.FileHeader
}

func NewFile(source multipart.File, header *multipart.FileHeader) File {
	return File{source: source, header: header}
}

func (f File) Source() multipart.File {
	return f.source
}

func (f File) Header() *multipart.FileHeader {
	return f.header
}

func (f File) Close() error {
	return f.source.Close()
}

func (f File) Content() string {
	var buff bytes.Buffer
	_, err := buff.ReadFrom(f.source)
	if err != nil {
		panic(err)
	}
	return buff.String()
}
