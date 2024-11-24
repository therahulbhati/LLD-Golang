package infrastructure

import (
	"fmt"
	"os"
	"sync"

	"logging-framework/domain"
)

type FileSink struct {
	mu   sync.Mutex
	file *os.File
}

func NewFileSink(filePath string) (*FileSink, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &FileSink{file: file}, nil
}

func (fs *FileSink) Log(message domain.Message) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	_, err := fs.file.WriteString(fmt.Sprintf("[%s] %s: %s\n",
		message.Level.String(), message.Namespace, message.Content))
	return err
}
