package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// ListDeadLetterSourceQueues is not implemented. It will panic in all cases.
func (client *SQS) ListDeadLetterSourceQueues(context.Context, *sqs.ListDeadLetterSourceQueuesInput, ...func(*sqs.Options)) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	panic("ListDeadLetterSourceQueues is not implemented")
}
