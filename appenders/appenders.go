/*
Package appenders contain interfaces and simple implementation for logger appenders
Logger use appenders interface to represent append mechanism for the logger events.
*/
package appenders

import (
	"github.com/yorikya/go-logger/encoders"
	"github.com/yorikya/go-logger/event"
)

// IAppender interface for append an incoming events.
type IAppender interface {
	//DoAppend append an event
	DoAppend(event.Event)
	//GetEncoder return encoder
	GetEncoder() encoders.IEncoder
}

// IAppenders interface for multiple appenders incoming event.
type IAppenders interface {
	//DoAppendAll append an event to multiple appenders.
	DoAppendAll(event.Event)
	//GetAppender return own appender
	GetAppender(int) IAppender
}

// Appenders struct implements IAppenders interface
type Appenders struct {
	//appenders holds items implements IAppender interface.
	appenders []IAppender
}

// NewAppenders return an new Appenders struct
// Args: As an argument NewAppenders get IAppender items
func NewAppenders(appenders ...IAppender) *Appenders {
	a := Appenders{}
	for _, apender := range appenders {
		a.appenders = append(a.appenders, apender)
	}
	return &a
}

// DoAppendAll interface implementation
func (a *Appenders) DoAppendAll(e event.Event) {
	for _, appender := range a.appenders {
		appender.DoAppend(e)
	}
}

//GetAppender return appender at particular index
func (a *Appenders) GetAppender(index int) IAppender {
	return a.appenders[index]
}
