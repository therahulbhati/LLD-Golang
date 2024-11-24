package infrastructure

import (
	"fmt"

	"logging-framework/domain"
)

type ConsoleSink struct{}

func (c *ConsoleSink) Log(message domain.Message) error {
	fmt.Printf("[%s] %s: %s\n", message.Level.String(), message.Namespace, message.Content)
	return nil
}
