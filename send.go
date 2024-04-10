package mocksqs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
)

// SendMessage is partially supported. The following are not supported:
//
// - SendMessageInput.DelaySeconds
//
// - SendMessageInput.MessageAttributes
//
// - SendMessageInput.MessageDeduplicationId
//
// - SendMessageInput.MessageGroupId
//
// - SendMessageInput.MessageSystemAttributes
//
// - SendMessageOutput.MD5OfMessageAttributes
//
// - SendMessageOutput.MD5OfMessageBody
//
// - SendMessageOutput.MD5OfMessageSystemAttributes
//
// - SendMessageOutput.MessageId
//
// - SendMessageOutput.SequenceNumber
func (client *SQS) sendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"SendMessageInput.QueueUrl":    input.QueueUrl,
		"SendMessageInput.MessageBody": input.MessageBody,
	})
	if err != nil {
		return &sqs.SendMessageOutput{}, err
	}

	if *input.MessageBody == "" {
		return &sqs.SendMessageOutput{}, errorMissingParameter("MessageBody")
	}

	if input.DelaySeconds < 0 || input.DelaySeconds > 900 {
		return &sqs.SendMessageOutput{}, errorInvalidParameterValue(
			fmt.Sprintf("Value %d for parameter DelaySeconds is invalid. Reason: DelaySeconds must be >= 0 and <= 900.",
				input.DelaySeconds))
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		receiptHandle := uuid.New().String()
		queue.messages.Set(receiptHandle, &Message{
			Message: types.Message{
				Body:          input.MessageBody,
				ReceiptHandle: aws.String(receiptHandle),
			},
			VisibleAfter: time.Now(),
		})

		return &sqs.SendMessageOutput{}, nil
	}

	return &sqs.SendMessageOutput{}, errorNonExistentQueue()
}

// SendMessageWithContext is partially supported. The following are not supported:
//
// - Recording ctx
//
// - Recording opts
//
// - Also see all features not supported for SendMessage()
func (client *SQS) SendMessage(ctx context.Context, input *sqs.SendMessageInput, _ ...func(*sqs.Options)) (*sqs.SendMessageOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	return client.sendMessage(input)
}

// SendMessageBatch is not implemented. It will panic in all cases.
func (client *SQS) SendMessageBatch(context.Context, *sqs.SendMessageBatchInput, ...func(*sqs.Options)) (*sqs.SendMessageBatchOutput, error) {
	panic("SendMessageBatch is not implemented")
}
