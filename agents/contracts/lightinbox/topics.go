package lightinbox

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedLightInbox, err := abi.JSON(strings.NewReader(LightInboxMetaData.ABI))
	if err != nil {
		panic(err)
	}

	AttestationAcceptedTopic = parsedLightInbox.Events["AttestationAccepted"].ID

	if AttestationAcceptedTopic == (common.Hash{}) {
		panic("AttestationAcceptedTopic is nil")
	}
}

// AttestationAcceptedTopic is the topic that gets emitted when the AttestationAccepted event is called.
var AttestationAcceptedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		AttestationAcceptedEvent: AttestationAcceptedTopic,
	}
}

// eventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func eventTypeFromTopic(ogTopic common.Hash) *EventType {
	fmt.Printf("eventTypeFromTopic on topic %s\n", ogTopic.Hex())
	for eventType, topic := range topicMap() {
		fmt.Printf("checking eventType %v and topic %v\n", eventType, topic.Hex())
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}
