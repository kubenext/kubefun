package errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGenericError(t *testing.T) {
	err := fmt.Errorf("access denied")

	intErr := NewGenericError(err)
	assert.Equal(t, "access denied", intErr.Error())
	assert.EqualError(t, intErr.err, "access denied")
	assert.NotEmpty(t, intErr.Timestamp())
	assert.NotZero(t, intErr.ID())
}
