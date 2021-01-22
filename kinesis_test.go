package kinesismanager_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/frozentech/kinesismanager"
)

type KinesisClient struct {
	kinesisiface.KinesisAPI
}

type TestObject struct {
	Name string `json:"name"`
}

// PutRecordWithContext impl
func (m *KinesisClient) PutRecordWithContext(ctx aws.Context, in *kinesis.PutRecordInput, args ...request.Option) (out *kinesis.PutRecordOutput, err error) {
	return &kinesis.PutRecordOutput{}, nil
}

func Test_Kinesis(t *testing.T) {
	ctx := context.WithValue(context.Background(), "x-amzn-trace-id", "Root=fakeid; Parent=reqid; Sampled=1")
	stream := kinesismanager.New("streamname")
	stream.Client = &KinesisClient{}
	stream.Publish(ctx, TestObject{}, "partition")
}
