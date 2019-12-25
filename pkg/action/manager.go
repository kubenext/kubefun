package action

import (
	"context"
	"github.com/kubenext/kubefun/internal/log"
	"sync"
	"time"
)

//go:generate mockgen -destination=./fake/mock_alert.go -package=fake github.com/kubenext/kubefun/pkg/action Alerter

const (
	// DefaultAlertExpiration is the default expiration for alerts.
	DefaultAlertExpiration = 10 * time.Second

	// AlertTypeError is for error alerts.
	AlertTypeError AlertType = "ERROR"

	// AlertTypeWarning is for warning alerts.
	AlertTypeWarning AlertType = "WARNING"

	// AlertTypeInfo is for info alerts.
	AlertTypeInfo AlertType = "INFO"
)

// AlertType is the type of alert.
type AlertType string

// Alert is an alert message.
type Alert struct {
	// Type is the type of alert.
	Type AlertType `json:"type"`
	// Message is the message for the alert.
	Message string `json:"message"`
	// Expiration is the time the alert expires.
	Expiration *time.Time `json:"expiration,omitempty"`
}

type Alerter interface {
	SendAlert(alert Alert)
}

// CreateAlert creates an alert with optional expiration. If the expireAt is < 1
// Expiration will be nil.
func CreateAlert(alertType AlertType, message string, expireAt time.Duration) Alert {
	alert := Alert{
		Type:    alertType,
		Message: message,
	}

	if expireAt > 0 {
		t := time.Now().Add(expireAt)
		alert.Expiration = &t
	}
	return alert
}

// DispatcherFunc is a function that will be dispatched to handle a payload.
type DispatcherFunc func(ctx context.Context, alerter Alerter, payload Payload) error

// Dispatcher handles actions.
type Dispatcher interface {
	ActionName() string
	Handle(ctx context.Context, alerter Alerter, payload Payload) error
}

// Dispatchers is a slice of Dispatcher.
type Dispatchers []Dispatcher

// ToActionPaths converts Dispatchers to a map.
func (d Dispatchers) ToActionPaths() map[string]DispatcherFunc {
	m := make(map[string]DispatcherFunc)

	for i := range d {
		m[d[i].ActionName()] = d[i].Handle
	}

	return m
}

// Manager manages actions.
type Manager struct {
	logger     log.Logger
	dispatches map[string]DispatcherFunc

	mu sync.Mutex
}

// NewManager creates an instance of Manager.
func NewManager(logger log.Logger) *Manager {
	return &Manager{
		logger:     logger.With("component", "action-manager"),
		dispatches: make(map[string]DispatcherFunc),
	}
}

// Register registers a dispatcher function to an action path.
func (m *Manager) Register(actionPath string, actionFunc DispatcherFunc) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.dispatches[actionPath] = actionFunc
	return nil
}

// Dispatch dispatches a payload to a path.
func (m *Manager) Dispatch(ctx context.Context, alerter Alerter, actionPath string, payload Payload) error {
	f, ok := m.dispatches[actionPath]
	if !ok {
		return &NotFoundError{Path: actionPath}
	}
	return f(ctx, alerter, payload)
}
