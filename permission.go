package mocksqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// AddPermission is not implemented. It will panic in all cases.
func (client *SQS) AddPermission(context.Context, *sqs.AddPermissionInput, ...func(*sqs.Options)) (*sqs.AddPermissionOutput, error) {
	panic("AddPermission is not implemented")
}

// RemovePermission is not implemented. It will panic in all cases.
func (client *SQS) RemovePermission(context.Context, *sqs.RemovePermissionInput, ...func(*sqs.Options)) (*sqs.RemovePermissionOutput, error) {
	panic("RemovePermission is not implemented")
}
