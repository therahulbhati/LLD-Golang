package infrastructure

import (
	"sync"

	"logging-framework/domain"
)

type InMemorySink struct {
	mu      sync.Mutex
	entries []domain.Message
}

func NewInMemorySink() *InMemorySink {
	return &InMemorySink{
		entries: make([]domain.Message, 0),
	}
}

func (ims *InMemorySink) Log(message domain.Message) error {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	ims.entries = append(ims.entries, message)
	return nil
}

func (ims *InMemorySink) GetEntries() []domain.Message {
	ims.mu.Lock()
	defer ims.mu.Unlock()
	return ims.entries
}
