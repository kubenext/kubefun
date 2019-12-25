package errors

import (
	"fmt"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewActionError(t *testing.T) {
	requestType := "errorList"
	payload := action.Payload{}
	err := fmt.Errorf("setNamespace error")

	intErr := NewActionError(requestType, payload, err)
	assert.Equal(t, payload, intErr.Payload())
	assert.Equal(t, requestType, intErr.RequestType())
	assert.Equal(t, fmt.Sprintf("%s: %s", requestType, err), intErr.Error())
	assert.EqualError(t, err, "setNamespace error")
	assert.NotEmpty(t, intErr.timestamp)
	assert.NotZero(t, intErr.id)
}
