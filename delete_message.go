package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// DeleteMessage is fully supported.
func (client *SQS) DeleteMessage(ctx context.Context, input *sqs.DeleteMessageInput, _ ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	return client.deleteMessage(input)
}

func (client *SQS) deleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"DeleteMessageInput.QueueUrl":      input.QueueUrl,
		"DeleteMessageInput.ReceiptHandle": input.ReceiptHandle,
	})
	if err != nil {
		return nil, err
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		didDelete := queue.delete(*input.ReceiptHandle)
		if !didDelete {
			return nil, errorInternal()
		}

		return nil, nil
	}

	return nil, errorNonExistentQueue()
}

// DeleteMessageBatch is fully supported.
func (client *SQS) DeleteMessageBatch(ctx context.Context, input *sqs.DeleteMessageBatchInput, _ ...func(*sqs.Options)) (*sqs.DeleteMessageBatchOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	for _, message := range input.Entries {
		_, err := client.deleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      input.QueueUrl,
			ReceiptHandle: message.ReceiptHandle,
		})

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
