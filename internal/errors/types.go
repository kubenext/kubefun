package errors

import "time"

// InternalError represents an internal Octant error.
type InternalError interface {
	ID() string
	Error() string
	Timestamp() time.Time
}
