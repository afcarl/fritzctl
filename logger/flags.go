package logger

import (
	"github.com/pkg/errors"
)

// Loglevel represents a Value for different logging configs.
type Loglevel struct {
}

// Type returns a name for the value type.
func (l *Loglevel) Type() string {
	return "string"
}

// String represents the default value.
func (l *Loglevel) String() string {
	return "info"
}

// Set configures the loglevel for the application.
func (l *Loglevel) Set(val string) error {
	err := configureLogLevel(val)
	if err != nil {
		return errors.Wrapf(err, "cannot apply loglevel configuration for value '%s'", val)
	}
	return nil
}
