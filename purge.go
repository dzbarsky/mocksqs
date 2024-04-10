package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// PurgeQueue is not implemented. It will panic in all cases.
func (client *SQS) PurgeQueue(context.Context, *sqs.PurgeQueueInput, ...func(*sqs.Options)) (*sqs.PurgeQueueOutput, error) {
	panic("PurgeQueue is not implemented")
}
