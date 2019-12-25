package errors

import (
	"fmt"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccessError(t *testing.T) {
	key := store.Key{
		Namespace:  "default",
		APIVersion: "v1",
		Kind:       "Pod",
	}
	verb := "watch"
	err := fmt.Errorf("access denied")

	intErr := NewAccessError(key, verb, err)
	assert.Equal(t, key, intErr.Key())
	assert.Equal(t, verb, intErr.Verb())
	assert.Equal(t, fmt.Sprintf("%s: %s: %s", verb, key, err), intErr.Error())
	assert.EqualError(t, err, "access denied")
	assert.NotEmpty(t, intErr.Timestamp())
	assert.NotZero(t, intErr.ID())
}
