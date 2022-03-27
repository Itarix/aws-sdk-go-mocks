package test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Itarix/aws-sdk-go-mocks/service/s3mock"
)

func TestLists3buckets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedOut := &s3.ListBucketsOutput{
		Owner: &s3.Owner{DisplayName: aws.String("some display name")},
	}
	m := s3mock.NewMockS3API(ctrl)
	m.EXPECT().
		ListBuckets(&s3.ListBucketsInput{}).
		Return(expectedOut, nil)

	out, err := listS3Buckets(m)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedOut, out)
	}
}
