package event

import "github.com/pkg/errors"

type notFound interface {
	NotFound() bool
	Path() string
}

var (
	errNotReady = errors.New("not ready")
)
