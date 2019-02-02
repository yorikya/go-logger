package appenders

import (
	"github.com/yorikya/go-logger/event"
)

type IAppender interface {
	DoAppend(event.Event)
}

type IAppenders interface {
	DoAppendAll(event.Event)
}

type Appenders struct {
	appenders []IAppender
}

func NewAppenders(appenders ...IAppender) Appenders {
	a := Appenders{}
	for _, apender := range appenders {
		a.appenders = append(a.appenders, apender)
	}
	return a
}

func (a Appenders) DoAppendAll(e event.Event) {
	for _, appender := range a.appenders {
		appender.DoAppend(e)
	}
}
