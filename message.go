package mocksqs

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Message struct {
	types.Message
	VisibleAfter time.Time
	ReceiveCount int64
}
