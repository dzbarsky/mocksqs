package mocksqs_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSQS_ReceiveMessage(t *testing.T) {
	t.Run("MissingQueueURL", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()

		_, err := client.client.ReceiveMessage(&sqs.ReceiveMessageInput{})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, ReceiveMessageInput.QueueUrl.\n")
	})

	t.Run("QueueDoesNotExist", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()

		_, err := client.client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: aws.String(mocksqs.CreateQueueURL("FOO")),
		})
		assertRegexpError(t, err, "AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.\n\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("ReceiveSingle", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()

		_, err := client.client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    &client.queueURL,
			MessageBody: aws.String("a"),
		})
		require.NoError(t, err)

		result, err := client.client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: aws.String(client.queueURL),
		})
		require.NoError(t, err)
		assert.Len(t, result.Messages, 1)
	})
}
