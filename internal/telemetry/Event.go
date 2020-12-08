package telemetry

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event is a telemetry event
type Event struct {
	ID     string
	Type   EventType
	UserID string
	Time   time.Time
	Data   map[DataKey]interface{}
}

var (
	executionID string
	user        string
)

// ConfigureEvents initializes any globals related to events{
func ConfigureEvents(pUser string) {
	executionID = primitive.NewObjectID().Hex()
	user = pUser
}

// NewCommandStartEvent creates a new event for a command's start
func NewCommandStartEvent(command string) *Event {
	data := make(map[DataKey]interface{})
	data[DataKeyCommand] = command
	return newEvent(EventTypeCommandStart, user, data)
}

// NewCommandCompleteEvent creates a new event for a command's start
func NewCommandCompleteEvent(command string) *Event {
	data := make(map[DataKey]interface{})
	data[DataKeyCommand] = command
	return newEvent(EventTypeCommandStart, user, data)
}

// NewCommandErrorEvent creates a new event for a command's start
func NewCommandErrorEvent(command string, err error) *Event {
	data := make(map[DataKey]interface{})
	data[DataKeyCommand] = command
	data[DataKeyErr] = err
	return newEvent(EventTypeCommandStart, user, data)
}

func newEvent(eventType EventType, user string, data map[DataKey]interface{}) *Event {
	data[DataKeyExecutionID] = executionID
	return &Event{
		ID:     primitive.NewObjectID().Hex(),
		Type:   eventType,
		UserID: user,
		Time:   time.Now(),
		Data:   data,
	}
}

// EventType is a cli event type
type EventType string

// set of supported cli event types
const (
	EventTypeCommandStart    EventType = "COMMAND_START"
	EventTypeCommandComplete EventType = "COMMAND_COMPLETE"
	EventTypeCommandError    EventType = "COMMAND_ERROR"
)

// DataKey used to pass data into the Event.Data map
type DataKey string

// set of Data Keys
const (
	DataKeyCommand     DataKey = "command"
	DataKeyExecutionID DataKey = "execution_id"
	DataKeyErr         DataKey = "err"
)
