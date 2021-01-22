package kinesismanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/kinesis"

	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

// KinesisManager ....
type KinesisManager struct {
	Name   string
	Client kinesisiface.KinesisAPI
}

// New ...
func New(name string) *KinesisManager {
	return &KinesisManager{
		Name:   name,
		Client: kinesis.New(session.New()),
	}
}

// Publish ...
func (me *KinesisManager) Publish(ctx context.Context, model interface{}, partition string, args ...string) (err error) {

	pl, _ := json.Marshal(model)

	input := &kinesis.PutRecordInput{
		StreamName:   aws.String(me.Name),
		PartitionKey: aws.String(partition),
		Data:         pl,
	}
	_, err = me.Client.PutRecordWithContext(ctx, input)

	return
}
