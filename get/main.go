package main

import (
	"context"
	"fmt"
	"io"
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

	bucketName := "test-bucket"     // SET THE BUCKET NAME
	objectKey := "my-file12345.txt" // SET THE FILE NAME

	file, err := os.Create(objectKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Get object content
	object, err := minioClient.GetObject(context.TODO(), bucketName, objectKey, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer object.Close()

	// Write object content to the file
	_, err = io.Copy(file, object)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Object content saved to:", objectKey)
	return
}
