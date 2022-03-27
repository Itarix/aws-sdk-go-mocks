// Package test purpose is just to make sure the project compiles
package test

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func listS3Buckets(api s3iface.S3API) (*s3.ListBucketsOutput, error) {
	in := &s3.ListBucketsInput{}
	return api.ListBuckets(in)
}
