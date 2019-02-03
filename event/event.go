package event

import (
	"sync"
	"time"

	"github.com/yorikya/go-logger/level"
)

const (
	missingStringKey = "????"
)

//eventPool relieving pressure on the garbage collector
var eventPool = sync.Pool{}

//Event alias point to relevant interface
type Event = ILogEvent

// ### Interfaces
type ILogEvent interface {
	GetMessage() string
	GetLevel() level.Level
	GetTimestamp() time.Time
	GetCaller() string
	GetLoggerName() string
	SetTimestamp(time.Time)
	SetCaller(string)
	SetLoggerName(string)
	ContainFlag(int) bool
}

// ### Implementations

type LogEvent struct {
	//Message string
	msg string
	//Message level tag
	lvl level.Level
	//Message Timestamp
	ts time.Time
	//Message caller
	caller string
	//Top level logger name
	loggerName string
	//Flags of root logger
	flags int
}

func NewBasicLogEvent(lvl level.Level, msg string, flags int) *LogEvent {
	var newEvt *LogEvent
	if poolEvent := eventPool.Get(); poolEvent != nil {
		newEvt = poolEvent.(*LogEvent)
	} else {
		newEvt = &LogEvent{}
	}
	// Basic log Event fields
	newEvt.setLevel(lvl)
	newEvt.setMessage(msg)
	newEvt.setFlags(flags)

	return newEvt
}

func ReleaseLogEvent(e *LogEvent) {
	eventPool.Put(e)
}

func (e *LogEvent) setMessage(msg string) {
	e.msg = msg
}

func (e *LogEvent) GetLoggerName() string {
	return e.loggerName
}

func (e *LogEvent) setLevel(lvl level.Level) {
	e.lvl = lvl
}

func (e *LogEvent) SetTimestamp(ts time.Time) {
	e.ts = ts
}

func (e *LogEvent) SetCaller(caller string) {
	e.caller = caller
}

func (e *LogEvent) setFlags(flags int) {
	e.flags = flags
}

func (e *LogEvent) SetLoggerName(name string) {
	e.loggerName = name
}

func (e *LogEvent) GetMessage() string {
	return e.msg
}

func (e *LogEvent) GetLevel() level.Level {
	return e.lvl
}

func (e *LogEvent) GetTimestamp() time.Time {
	return e.ts
}

func (e *LogEvent) GetCaller() string {
	return e.caller
}

func (e *LogEvent) GetFlags() int {
	return e.flags
}

func (e *LogEvent) ContainFlag(flag int) bool {
	return e.flags&flag != 0
}
