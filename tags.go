package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// ListQueueTags is not implemented. It will panic in all cases.
func (client *SQS) ListQueueTags(context.Context, *sqs.ListQueueTagsInput, ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error) {
	panic("ListQueueTags is not implemented")
}

// TagQueue is not implemented. It will panic in all cases.
func (client *SQS) TagQueue(context.Context, *sqs.TagQueueInput, ...func(*sqs.Options)) (*sqs.TagQueueOutput, error) {
	panic("TagQueue is not implemented")
}

// UntagQueue is not implemented. It will panic in all cases.
func (client *SQS) UntagQueue(context.Context, *sqs.UntagQueueInput, ...func(*sqs.Options)) (*sqs.UntagQueueOutput, error) {
	panic("UntagQueue is not implemented")
}
