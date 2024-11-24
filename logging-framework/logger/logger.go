package logger

import (
	"errors"
	"time"

	"logging-framework/config"
	"logging-framework/domain"
	"logging-framework/infrastructure"
)

type Logger struct {
	sinks map[domain.Level]domain.Sink
}

func NewLogger() *Logger {
	return &Logger{
		sinks: make(map[domain.Level]domain.Sink),
	}
}

func ConfigureLogger(cfg *config.Config) (*Logger, error) {
	logr := NewLogger()
	for levelStr, sinkType := range cfg.LogLevels {
		sinkCfg := infrastructure.SinkConfig{
			SinkType: sinkType,
			FilePath: cfg.FilePath,
			DSN:      cfg.DSN,
		}
		sink := infrastructure.CreateSink(sinkCfg)

		if sink == nil {
			return nil, errors.New("invalid sink type")
		}

		if level, err := domain.ParseLevel(levelStr); err == nil {
			logr.Configure(level, sink)
		} else {
			return nil, err
		}

	}
	return logr, nil
}

func (l *Logger) Configure(level domain.Level, sink domain.Sink) {
	l.sinks[level] = sink
}

func (l *Logger) Log(content string, level domain.Level, namespace string) error {
	message := domain.Message{
		Content:   time.Now().Format(time.RFC3339) + " " + content,
		Level:     level,
		Namespace: namespace,
	}

	for lvl, sink := range l.sinks {
		if lvl >= level {
			if err := sink.Log(message); err != nil {
				return err
			}
		}
	}
	return nil
}
