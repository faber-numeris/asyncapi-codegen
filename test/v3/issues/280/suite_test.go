//go:generate go run ../../../../cmd/asyncapi-codegen -p issue280 -i ./asyncapi.yaml -o ./asyncapi.gen.go

package issue280

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, NewSuite())
}

type Suite struct {
	suite.Suite
}

func NewSuite() *Suite {
	return &Suite{}
}

func (suite *Suite) TestHeadersRefResolution() {
	// Verify that the KafkaHeadersSchema struct is generated correctly
	// with the expected fields from the referenced schema
	var headers KafkaHeadersSchema
	assert.NotNil(suite.T(), headers)

	// Verify TestMessageMessageFromTestingChannel has Headers field of correct type
	var msg TestMessageMessageFromTestingChannel
	correlationID := "test-correlation-id"
	msg.Headers.KafkaCorrelationId = &correlationID

	assert.Equal(suite.T(), "test-correlation-id", *msg.Headers.KafkaCorrelationId)
}

func (suite *Suite) TestToBrokerMessageWithHeaders() {
	// Create a message with headers
	msg := TestMessageMessageFromTestingChannel{
		Headers: KafkaHeadersSchema{
			KafkaCorrelationId: ptrString("test-correlation-id"),
			KafkaKey:           ptrString("test-key"),
		},
		Payload: TestPayloadSchema{
			Message: ptrString("test message"),
		},
	}

	// Convert to broker message
	brokerMsg, err := msg.toBrokerMessage()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), brokerMsg.Headers)
	assert.Contains(suite.T(), brokerMsg.Headers, "kafka_correlationId")
	assert.Contains(suite.T(), brokerMsg.Headers, "kafka_key")
}

func ptrString(s string) *string {
	return &s
}
