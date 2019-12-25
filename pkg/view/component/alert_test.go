package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlert(t *testing.T) {
	got := NewAlert(AlertTypeSuccess, "message")
	expected := Alert{
		Type:    AlertTypeSuccess,
		Message: "message",
	}

	assert.Equal(t, got, expected)
}
