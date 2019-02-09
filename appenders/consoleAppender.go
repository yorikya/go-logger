/*
Package appenders | ConsoleAppender
ConsoleAppender write logger event to console, with given encoder style.
*/
package appenders

import (
	"github.com/yorikya/go-logger/encoders"
	"github.com/yorikya/go-logger/event"
)

//ConsoleAppender implements Appender interface
type ConsoleAppender struct {
	//encoder console encoder
	encoder encoders.IEncoder
}

//NewConsoleAppender return a new ConsoleAppender with encoder.
func NewConsoleAppender(enc encoders.IEncoder) *ConsoleAppender {
	c := ConsoleAppender{
		encoder: enc,
	}
	return &c
}

//DoAppend encode an incoming event with own encoder.
func (a *ConsoleAppender) DoAppend(e event.Event) {
	a.encoder.Encode(e)
}
