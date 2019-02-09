/*
Package event contain interfaces and simple implementation for logger events
Logger create an new event from each incoming message, represnet event behavior .
*/
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

// ILogEvent event interface
type ILogEvent interface {
	// GetMessage return event (Type:string)
	GetMessage() string
	// GetLevel return event level (Type:level.Level)
	GetLevel() level.Level
	// GetTimestamp return event timestamp (Type:Time)
	GetTimestamp() time.Time
	// GetCaller return event caller (Type:string)
	GetCaller() string
	//GetLoggerName return event logger name
	GetLoggerName() string
	//SetTimestamp set event creation timestamp
	SetTimestamp(time.Time)
	//SetCaller set event caller
	SetCaller(string)
	//SetLoggerName set event logger name
	SetLoggerName(string)
	//ContainFlag return true if an event contain particular Flag
	ContainFlag(int) bool
}

// LogEvent implements IEvent interface
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

//NewBasicLogEvent return new LogEvent with basic fields given as an argument
func NewBasicLogEvent(lvl level.Level, msg string, flags int) *LogEvent {
	var newEvt *LogEvent
	if poolEvent := eventPool.Get(); poolEvent != nil { //get new LogEvent from pool
		newEvt = poolEvent.(*LogEvent)
	} else { // create a new LogEvent
		newEvt = &LogEvent{}
	}
	// Basic log Event fields
	newEvt.setLevel(lvl)
	newEvt.setMessage(msg)
	newEvt.setFlags(flags)

	return newEvt
}

// ReleaseLogEvent release event to the pool
func ReleaseLogEvent(e *LogEvent) {
	eventPool.Put(e)
}

//setMessage interface function implementation
func (e *LogEvent) setMessage(msg string) {
	e.msg = msg
}

// GetLoggerName interface function implementation
func (e *LogEvent) GetLoggerName() string {
	return e.loggerName
}

// setLevel interface function implementation
func (e *LogEvent) setLevel(lvl level.Level) {
	e.lvl = lvl
}

// SetTimestamp interface function implementation
func (e *LogEvent) SetTimestamp(ts time.Time) {
	e.ts = ts
}

// SetCaller interface function implementation
func (e *LogEvent) SetCaller(caller string) {
	e.caller = caller
}

// setFlags set flags to event
func (e *LogEvent) setFlags(flags int) {
	e.flags = flags
}

// SetLoggerName interface function implementation
func (e *LogEvent) SetLoggerName(name string) {
	e.loggerName = name
}

// GetMessage interface function implementation
func (e *LogEvent) GetMessage() string {
	return e.msg
}

// GetLevel interface function implementation
func (e *LogEvent) GetLevel() level.Level {
	return e.lvl
}

// GetTimestamp interface function implementation
func (e *LogEvent) GetTimestamp() time.Time {
	return e.ts
}

// GetCaller interface function implementation
func (e *LogEvent) GetCaller() string {
	return e.caller
}

// GetFlags interface function implementation
func (e *LogEvent) GetFlags() int {
	return e.flags
}

// ContainFlag interface function implementation
func (e *LogEvent) ContainFlag(flag int) bool {
	return e.flags&flag != 0
}
