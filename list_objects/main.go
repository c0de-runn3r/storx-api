package main

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	accessKeyID := os.Getenv("ACCESS_KEY")
	if accessKeyID == "" {
		fmt.Println("ACCESS_KEY is not set!")
		return
	}
	secretAccessKey := os.Getenv("SECRET_KEY")
	if secretAccessKey == "" {
		fmt.Println("SECRET_KEY is not set!")
		return
	}
	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		fmt.Println("ENDPOINT is not set!")
		return
	}
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bucketName := "test-bucket" // SET THE BUCKET NAME

	objectsCh := minioClient.ListObjects(context.TODO(), bucketName, minio.ListObjectsOptions{})
	var objects []minio.ObjectInfo
	for object := range objectsCh {
		if object.Err != nil {
			return
		}
		objects = append(objects, object)
	}
	fmt.Println("Lis of objects in bucket ", bucketName, ": ", objects)

}
