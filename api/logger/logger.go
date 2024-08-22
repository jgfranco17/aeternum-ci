package logger

import (
	"context"
	"strings"

	"api/context_settings"

	"github.com/sirupsen/logrus"
)

var stringToLogLevel map[string]logrus.Level

func init() {

	stringToLogLevel = map[string]logrus.Level{
		"DEBUG": logrus.DebugLevel,
		"INFO":  logrus.InfoLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
		"PANIC": logrus.PanicLevel,
		"FATAL": logrus.FatalLevel,
		"TRACE": logrus.TraceLevel,
	}
}

// Returns an instance of the logger adding the fields found in the context.
func FromContext(ctx context.Context) *logrus.Entry {
	entry := logrus.WithFields(logrus.Fields{})

	if ctx == nil {
		return entry
	}

	// List of fields we want to expose as part of the logged messages
	fields := []string{
		context_settings.RequestId,
		context_settings.Version,
		context_settings.Environment,
		context_settings.Origin,
	}

	for _, field := range fields {
		value := ctx.Value(field)

		if value != nil {
			entry = entry.WithField(string(field), value)
		}
	}

	return entry
}

func SetLevel(level string) {
	loglevel, ok := stringToLogLevel[strings.ToUpper(level)]

	if ok {
		logrus.SetLevel(loglevel)
	}

}
