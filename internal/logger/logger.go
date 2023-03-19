package logger

import (
	"backend/internal/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
	"io"
	"os"
	"reflect"
)

type Logger struct {
	*logrus.Logger
}

func (l *Logger) LogEvent(event fxevent.Event) {
	log := l.WithField("module", "fx")
	typeName := reflect.TypeOf(event)
	log.Debugf("event: %s", typeName)
}

func NewLogger() *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(new(logrus.JSONFormatter))
	l.SetLevel(logrus.DebugLevel)

	return &Logger{Logger: l}
}

func NewProvider(cfg *config.Config) *Logger {
	l := NewLogger()

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		l.Fatal(err)
	}
	l.SetLevel(level)

	return l
}

func NewProviderWithDiscardOutput() *Logger {
	l := NewLogger()
	l.SetOutput(io.Discard)
	return l
}
