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

	buckets, err := minioClient.ListBuckets(context.TODO())
	if err != nil {
		return
	}
	fmt.Println("List of buckets:", buckets)

}
