package main

import (
	"fmt"
	"log"
	"os"

	"github.com/muzfr7/aws-s3-bucket-creator/config"

	"github.com/kelseyhightower/envconfig"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var env config.EnvConfig

func init() {
	// export environment variables from .env file
	if _, err := os.Stat("./.env"); err == nil {
		if err = config.ExportEnvVars("./.env"); err != nil {
			log.Fatal(err)
		}
	}

	// load env vars into EnvConfig struct
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
	}
}

func main() {

	// get bucket name
	var bucketName string
	fmt.Print("Enter new bucket name: ")
	fmt.Scan(&bucketName)

	// check whether bucket name was provided as an argument
	if len(bucketName) < 5 {
		log.Fatal("Bucket name must be provided!")
	}

	// start a new aws session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(env.AWSRegion),
		Credentials: credentials.NewSharedCredentials("", env.AWSProfile),
	})
	if err != nil {
		log.Fatal(err)
	}

	// create a new instance of aws service
	svc := s3.New(sess)

	// create an s3 bucket
	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Unable to create bucket %q, %v", bucketName, err)
	}

	log.Println("Waiting for bucket creation..")

	// wait until bucket is created before finishing
	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Error occurred while waiting for bucket to be created, %v", bucketName)
	}

	log.Println("Bucket created successfully..")
}
