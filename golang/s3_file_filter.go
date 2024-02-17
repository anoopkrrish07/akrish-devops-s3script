package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the s3 bucket name:")
	bucketName, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading bucket name: %v", err)
	}
	bucketName = strings.TrimSpace(bucketName)
	fmt.Println("Enter the search string:")
	searchString, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading search string: %v", err)
	}
	searchString = strings.TrimSpace(searchString)
	fmt.Println("Enter the Buckets AWS region:")
	region, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error reading AWS region: %v", err)
	}
	region = strings.TrimSpace(region)
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
	)

	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	svc := s3.NewFromConfig(cfg)
	fmt.Printf("Initiating search for files containing '%s'...\n", searchString)
	paginator := s3.NewListObjectsV2Paginator(svc, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: nil,
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			log.Printf("Failed to retrieve page: %v\n", err)
			continue
		}

		for _, obj := range page.Contents {
			if strings.HasSuffix(*obj.Key, ".txt") {
				result, err := svc.GetObject(context.Background(), &s3.GetObjectInput{
					Bucket: aws.String(bucketName),
					Key:    obj.Key,
				})

				if err != nil {
					log.Printf("Error downloading '%s': %v\n", *obj.Key, err)
					continue
				}
				buf := new(bytes.Buffer)
				_, err = buf.ReadFrom(result.Body)
				result.Body.Close()

				if err != nil {
					log.Printf("Error reading content of '%s': %v\n", *obj.Key, err)
					continue
				}

				if strings.Contains(buf.String(), searchString) {
					fmt.Printf("Match found in: %s\n", *obj.Key)
				}
			}
		}
	}
}
