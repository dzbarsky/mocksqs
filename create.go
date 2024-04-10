package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// CreateQueue is partially implemented. The following fields are not
// implemented:
//
// - Attributes: DelaySeconds, MaximumMessageSize, MessageRetentionPeriod,
// Policy, ReceiveMessageWaitTimeSeconds, RedrivePolicy, VisibilityTimeout,
// KmsMasterKeyId, KmsDataKeyReusePeriodSeconds, FifoQueue,
// ContentBasedDeduplication
//
// - Tags
func (client *SQS) CreateQueue(ctx context.Context, input *sqs.CreateQueueInput, _ ...func(*sqs.Options)) (*sqs.CreateQueueOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	if input.QueueName == nil {
		return &sqs.CreateQueueOutput{},
			errorMissingField("CreateQueueInput.QueueName")
	}

	queueURL := CreateQueueURL(*input.QueueName)

	if client.GetQueue(queueURL) == nil {
		client.queues.Store(queueURL, newQueue(queueURL))
	}

	return &sqs.CreateQueueOutput{
		QueueUrl: &queueURL,
	}, nil
}
