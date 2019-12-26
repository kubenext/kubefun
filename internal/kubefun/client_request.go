package kubefun

import "github.com/kubenext/kubefun/pkg/action"

// ClientRequestHandler is a client request.
type ClientRequestHandler struct {
	RequestType string
	Handler     func(state State, payload action.Payload) error
}
