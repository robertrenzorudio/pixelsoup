package awsqueue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type AWSQueue struct {
	client            *sqs.Client
	queueName         string
	queueURL          *string
	waitTimeSeconds   int32
	visibilityTimeout int32
}

func New(ctx context.Context, queueName string) (*AWSQueue, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-1"))

	if err != nil {
		return nil, err
	}

	q := AWSQueue{
		client: sqs.NewFromConfig(cfg), queueName: queueName}

	err = q.getQueueUrl(ctx)
	if err != nil {
		return nil, err
	}

	q.waitTimeSeconds = 20
	q.visibilityTimeout = 30
	return &q, nil

}

func (q *AWSQueue) getQueueUrl(ctx context.Context) error {
	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(q.queueName),
	}

	res, err := q.client.GetQueueUrl(ctx, input)
	if err != nil {
		return err
	}

	q.queueURL = res.QueueUrl
	return nil
}

func (q *AWSQueue) ReceiveMessage(ctx context.Context) (*sqs.ReceiveMessageOutput, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl: q.queueURL,
		AttributeNames: []types.QueueAttributeName{
			"SentTimestamp",
			"MessageGroupId",
		},
		MaxNumberOfMessages: 10,
		MessageAttributeNames: []string{
			"All",
		},
		WaitTimeSeconds:   q.waitTimeSeconds,
		VisibilityTimeout: q.visibilityTimeout,
	}

	jobs, err := q.client.ReceiveMessage(ctx, input)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (q *AWSQueue) DeleteMessageBatch(ctx context.Context, entries []types.DeleteMessageBatchRequestEntry) (*sqs.DeleteMessageBatchOutput, error) {
	removeMsgInput := &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: q.queueURL,
	}

	res, err := q.client.DeleteMessageBatch(context.TODO(), removeMsgInput)

	if err != nil {
		return nil, err
	}

	return res, nil
}
