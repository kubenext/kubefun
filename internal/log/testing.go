package log

import "testing"

type testingLogger struct {
	t *testing.T
}

// TestLogger returns a logger for tests
func TestLogger(t *testing.T) Logger {
	return &testingLogger{t: t}
}

func (t *testingLogger) Debugf(template string, args ...interface{}) {
	t.t.Logf(template, args...)
}

func (t *testingLogger) Infof(template string, args ...interface{}) {
	t.t.Logf(template, args...)
}

func (t *testingLogger) Warnf(template string, args ...interface{}) {
	t.t.Logf(template, args...)
}

func (t *testingLogger) Errorf(template string, args ...interface{}) {
	t.t.Errorf(template, args...)
}

func (t *testingLogger) With(args ...interface{}) Logger {
	return t
}

func (t *testingLogger) WithErr(err error) Logger {
	return t
}

func (t *testingLogger) Named(name string) Logger {
	return t
}
