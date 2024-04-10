package mocksqs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (client *SQS) changeMessageVisibility(input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		if message, ok := queue.messages.Get(*input.ReceiptHandle); ok {
			message.(*Message).VisibleAfter = time.Now().Add(
				time.Duration(input.VisibilityTimeout) * time.Second)
		}

		return nil, nil
	}

	return nil, fmt.Errorf("no such queue: %s", *input.QueueUrl)
}

func (client *SQS) ChangeMessageVisibility(ctx context.Context, input *sqs.ChangeMessageVisibilityInput, _ ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	return client.changeMessageVisibility(input)
}

// ChangeMessageVisibilityBatch is not implemented. It will panic in all cases.
func (client *SQS) ChangeMessageVisibilityBatch(context.Context, *sqs.ChangeMessageVisibilityBatchInput, ...func(*sqs.Options)) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	panic("ChangeMessageVisibilityBatch is not implemented")
}
