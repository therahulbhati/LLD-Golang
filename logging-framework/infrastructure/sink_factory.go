package infrastructure

import (
	"logging-framework/config"
	"logging-framework/domain"
)

type SinkConfig struct {
	SinkType string
	FilePath string
	DSN      string
}

func CreateSink(cfg SinkConfig) domain.Sink {
	switch cfg.SinkType {
	case config.ConsoleSinkType:
		return &ConsoleSink{}
	case config.InMemorySinkType:
		return NewInMemorySink()
	case config.FileSinkType:
		sink, err := NewFileSink(cfg.FilePath)
		if err != nil {
			return nil
		}
		return sink
	default:
		return nil
	}

}
