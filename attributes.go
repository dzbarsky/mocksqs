package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// GetQueueAttributes is not implemented. It will panic in all cases.
func (client *SQS) GetQueueAttributes(context.Context, *sqs.GetQueueAttributesInput, ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error) {
	panic("GetQueueAttributes is not implemented")
}

// SetQueueAttributes is not implemented. It will panic in all cases.
func (client *SQS) SetQueueAttributes(context.Context, *sqs.SetQueueAttributesInput, ...func(*sqs.Options)) (*sqs.SetQueueAttributesOutput, error) {
	panic("SetQueueAttributes is not implemented")
}
