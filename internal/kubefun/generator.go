package kubefun

import (
	"context"
	"time"
)

//go:generate mockgen -destination=./fake/mock_generator.go -package=fake github.com/kubenext/kubefun/internal/kubefun Generator

// Generator generates events.
type Generator interface {
	// Event generates events using the returned channel.
	Event(ctx context.Context) (Event, error)

	// ScheduleDelay is how long to wait before scheduling this generator again.
	ScheduleDelay() time.Duration

	// Name is the generator name.
	Name() string
}
