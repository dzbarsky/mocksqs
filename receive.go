package mocksqs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// ReceiveMessage is partially supported. The following are not supported:
//
// - ReceiveMessageInput.AttributeNames
//
// - ReceiveMessageInput.MessageAttributeNames
//
// - ReceiveMessageInput.ReceiveRequestAttemptId
//
// - ReceiveMessageInput.VisibilityTimeout
//
// - ReceiveMessageInput.WaitTimeSeconds
func (client *SQS) ReceiveMessage(ctx context.Context, input *sqs.ReceiveMessageInput, _ ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	if input.QueueUrl == nil {
		return nil, errorMissingField("ReceiveMessageInput.QueueUrl")
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		output := &sqs.ReceiveMessageOutput{}

		for el := queue.messages.Front(); el != nil; el = el.Next() {
			message := el.Value.(*Message)
			t := time.Now()
			if t.After(message.VisibleAfter) || t == message.VisibleAfter {
				timeout := input.VisibilityTimeout
				if timeout == 0 {
					timeout = 30
				}
				_, _ = client.changeMessageVisibility(&sqs.ChangeMessageVisibilityInput{
					QueueUrl:          input.QueueUrl,
					ReceiptHandle:     message.ReceiptHandle,
					VisibilityTimeout: timeout,
				})

				message.ReceiveCount++

				output.Messages = append(output.Messages, types.Message{
					Body:          message.Body,
					ReceiptHandle: message.ReceiptHandle,
					Attributes: map[string]string{
						"ApproximateReceiveCount": fmt.Sprintf("%d", message.ReceiveCount),
					},
				})

				if input.MaxNumberOfMessages == 0 {
					input.MaxNumberOfMessages = 1
				}

				if len(output.Messages) == int(input.MaxNumberOfMessages) {
					return output, nil
				}
			}
		}

		return output, nil
	}

	return nil, errorNonExistentQueue()
}
