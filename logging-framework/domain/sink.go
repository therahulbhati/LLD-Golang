package domain

type Sink interface {
	Log(message Message) error
}
