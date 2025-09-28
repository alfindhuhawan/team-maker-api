package util

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"team-maker-api/shared/infrastructure/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nfnt/resize"
)

func UploadImageGCS(fileUpload *multipart.FileHeader, AWSConfig config.GCS) (string, error) {

	var (
		gcsPath string
		err     error
	)

	src, err := fileUpload.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// CONVERT
	// Decoding gives you an Image.
	// If you have an io.Reader already, you can give that to Decode
	// without reading it into a []byte.
	img, _, err := image.Decode(src)
	// check err
	if err != nil {
		return "error converting", err
	}

	newImage := resize.Resize(160, 160, img, resize.Lanczos3)

	fileName := GenerateID(16) + ".jpg"

	// Encode uses a Writer, use a Buffer if you need the raw []byte
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, newImage, &jpeg.Options{Quality: 70})
	// check err
	if err != nil {
		return "", err
	}
	buffer := buf.Bytes()

	//// begin AWS
	creds := credentials.NewStaticCredentials(AWSConfig.GCSKey, AWSConfig.GCSSecret, "")
	_, err = creds.Get()
	if err != nil {
		return "", err
	}

	cfg := aws.NewConfig().WithRegion(AWSConfig.GCSRegion).WithCredentials(creds).WithEndpoint(AWSConfig.GCSUrl)
	svc := s3.New(session.New(), cfg)

	// DIRECTORY NAME BASED ON Y/M/D
	var directoryName string
	now := time.Now()
	directoryName = now.Format("2006/02/01")

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	gcsPath = directoryName + "/" + fileName
	params := &s3.PutObjectInput{
		Bucket: aws.String(AWSConfig.GCSBucket),
		Key:    aws.String(gcsPath),
		Body:   fileBytes,
		// ContentLength: aws.Int64(size), // TODO : TEMPORARY DISABLED
		ContentType: aws.String(fileType),
	}

	_, err = svc.PutObject(params)
	if err != nil {
		return "", err
	}

	return gcsPath, err
}
