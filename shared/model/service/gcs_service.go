package service

import (
	"bytes"
	"context"
)

type UploadCsvGCSService interface {
	UploadCsvGCS(ctx context.Context, csvBytes []byte, csvFileName string) (string, error)
}

type DirectUploadFileGCSService interface {
	DirectUploadFileGCS(ctx context.Context, filename string, csvBuffer *bytes.Buffer) (*DirectUploadRes, error)
}

type DirectUploadRes struct {
	Name string
}

type GetSignatureURLService interface {
	GetSignatureURL(ctx context.Context, fileName string) (string, error)
}
