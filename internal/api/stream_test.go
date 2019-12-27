package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_notFoundRedirectPath(t *testing.T) {
	cases := []struct {
		name     string
		expected string
	}{
		{
			name:     "overview/namespace/default/workloads/deployments/nginx-deployment/",
			expected: "overview/namespace/default/workloads/deployments",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := notFoundRedirectPath(tc.name)
			assert.Equal(t, tc.expected, got)
		})
	}
}
