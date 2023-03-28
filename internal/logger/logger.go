package logger

import (
	"io"
	"os"
	"reflect"

	"backend/internal/config"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
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
	log := NewLogger()

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(level)

	return log
}

func NewProviderWithDiscardOutput() *Logger {
	l := NewLogger()
	l.SetOutput(io.Discard)

	return l
}
