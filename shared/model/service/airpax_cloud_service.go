package service

import (
	"bytes"
	"context"
)

type UploadToAirpaxCloudService interface {
	UploadToAirpaxCloud(ctx context.Context, fileName string, csvBuffer *bytes.Buffer) (*FileInfo, error)
}

type FileInfo struct {
	// ID              vo.FileInfoID `json:"id,omitempty" bson:"_id"`                  //
	Name string `json:"name" bson:"name"` //
	// Extension       string        `json:"extension" bson:"extension"`               //
	// Path            string        `json:"path" bson:"path"`                         //
	// Folder          string        `json:"folder" bson:"folder"`                     //
	// BaseURL         string        `json:"base_url" bson:"base_url"`                 //
	// BucketName      string        `json:"bucket_name" bson:"bucket_name"`           //
	// StorageProvider string        `json:"storage_provider" bson:"storage_provider"` //
	// FullPath        string        `json:"full_path" bson:"full_path"`               //
	// CreatedAt       time.Time     `json:"created_at" bson:"created_at"`             //
	// UpdatedAt       time.Time     `json:"updated_at" bson:"updated_at"`             //
	SignatureURL string `json:"signature_url"`
}

type AirpaxCloudUploadResponse struct {
	Success      bool      `json:"success"`
	ErrorCode    string    `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
	Data         *FileInfo `json:"data"`
}

type DataDownloadResponse struct {
	Item []byte `json:"item"`
}

type AirpaxCloudDownloadResponse struct {
	Success      bool                  `json:"success"`
	ErrorCode    string                `json:"errorCode"`
	ErrorMessage string                `json:"errorMessage"`
	Data         *DataDownloadResponse `json:"data"`
}

type DownloadFileByNameService interface {
	DownloadFileByName(ctx context.Context, fileName string) ([]byte, error)
}
