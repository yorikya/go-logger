package appenders

import (
	"github.com/yorikya/go-logger/encoders"
	"github.com/yorikya/go-logger/event"
)

type ConsoleAppender struct {
	encoder encoders.Encoder
}

func NewConsoleAppender(enc encoders.Encoder) *ConsoleAppender {
	c := ConsoleAppender{
		encoder: enc,
	}
	return &c
}

func (a *ConsoleAppender) DoAppend(e event.Event) {
	a.encoder.Encode(e)
}
