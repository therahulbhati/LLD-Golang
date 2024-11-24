package util

import (
	"fmt"

	"github.com/google/uuid"
)

// GenerateUUIDWithPrefix generates a UUID with the given prefix.
func GenerateUUIDWithPrefix(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.New().String())
}
