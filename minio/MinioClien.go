package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

// document: https://www.minio.org.cn/docs/minio/linux/developers/go/minio-go.html

var (
	endpoint     string = "hadoop101:9000"
	accessKey    string = "minioadmin"
	accessSecret string = "minioadmin"
)

func main() {

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	// list bucket
	buckets, err := client.ListBuckets(ctx)
	if err != nil {
		fmt.Println("获取 Minio 桶出错...", err)
		panic(err)
	}
	for _, bucket := range buckets {
		fmt.Println(bucket.Name)
	}

	// new bucket
	bucketName := "go-minio"
	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, err := client.BucketExists(ctx, bucketName)
		if err == nil && exists {
			fmt.Printf("创建 bucket [%s] 已经存在\n", bucketName)
		} else {
			panic(err)
		}
	}

	// upload file
	filePath := "D:\\111.txt"
	info, err := client.FPutObject(ctx, bucketName, "go.txt", filePath, minio.PutObjectOptions{})
	if err != nil {
		fmt.Printf("上传文件 %s 到桶 %s 出错\n", filePath, bucketName)
		panic(err)
	}
	fmt.Println(info)

	// download file
	downloadFilePath := "D:\\112.txt"
	err = client.FGetObject(ctx, bucketName, "go.txt", downloadFilePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Printf("从桶 %s 下载文件 go.txt 出错\n", bucketName)
	}
}
