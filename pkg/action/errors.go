package action

import "fmt"

type NotFoundError struct {
	Path string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("action path %q not found", e.Path)
}
