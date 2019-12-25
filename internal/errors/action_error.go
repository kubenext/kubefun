package errors

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kubenext/kubefun/pkg/action"
	"time"
)

type ActionError struct {
	id          string
	timestamp   time.Time
	payload     action.Payload
	requestType string
	err         error
}

var _ InternalError = (*ActionError)(nil)

func NewActionError(requestType string, payload action.Payload, err error) *ActionError {
	id, _ := uuid.NewUUID()

	return &ActionError{
		requestType: requestType,
		payload:     payload,
		err:         err,
		timestamp:   time.Now(),
		id:          id.String(),
	}
}

// ID returns the error unique ID.
func (o *ActionError) ID() string {
	return o.id
}

// Timestamp returns the error timestamp.
func (o *ActionError) Timestamp() time.Time {
	return o.timestamp
}

// Error returns an error string.
func (o *ActionError) Error() string {
	return fmt.Sprintf("%s: %s", o.requestType, o.err)
}

// Client returns a client if one is available.
func (o *ActionError) RequestType() string {
	return o.requestType
}

// Request returns the payload that generated the error, if available.
func (o *ActionError) Payload() action.Payload {
	return o.payload
}
