package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// DeleteQueue is fully supported.
func (client *SQS) DeleteQueue(ctx context.Context, input *sqs.DeleteQueueInput, _ ...func(*sqs.Options)) (*sqs.DeleteQueueOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	if input.QueueUrl == nil {
		return &sqs.DeleteQueueOutput{},
			errorMissingField("DeleteQueueInput.QueueUrl")
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		client.queues.Delete(*input.QueueUrl)

		return &sqs.DeleteQueueOutput{}, nil
	}

	return &sqs.DeleteQueueOutput{},
		errorWithRequestID("AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.")
}
