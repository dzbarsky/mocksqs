package mocksqs

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// ListQueues is fully supported.
func (client *SQS) ListQueues(ctx context.Context, input *sqs.ListQueuesInput, _ ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error) {
	client.httpRequest()

	prefix := ""
	if input.QueueNamePrefix != nil {
		prefix = *input.QueueNamePrefix
	}

	output := &sqs.ListQueuesOutput{}

	client.queues.Range(func(key, value interface{}) bool {
		if strings.HasPrefix(key.(string), prefix) {
			output.QueueUrls = append(output.QueueUrls, value.(*Queue).URL)
		}

		return true
	})

	return output, nil
}
