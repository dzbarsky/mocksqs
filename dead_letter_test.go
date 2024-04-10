package mocksqs_test

import (
	"testing"

	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
)

func TestSQS_ListDeadLetterSourceQueues(t *testing.T) {
	assert.PanicsWithValue(t, "ListDeadLetterSourceQueues is not implemented", func() {
		mocksqs.New().ListDeadLetterSourceQueues(nil)
	})
}
