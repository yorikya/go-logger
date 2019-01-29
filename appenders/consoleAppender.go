package appenders

import (
	"gitlab.appsflyer.com/Architecture/af-logger-go/encoders"
	"gitlab.appsflyer.com/Architecture/af-logger-go/event"
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
