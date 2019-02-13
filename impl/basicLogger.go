//Should be moved to independnt module.
/*
Package impl holds implementation for Logger interface
*/
package impl

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/yorikya/go-logger/appenders"
	"github.com/yorikya/go-logger/encoders"
	"github.com/yorikya/go-logger/event"
	"github.com/yorikya/go-logger/filters"
	"github.com/yorikya/go-logger/flags"
	"github.com/yorikya/go-logger/level"
)

const (
	newLine string = "%s\n"
	//FBasicLoggerFlags basic logger default flags
	FBasicLoggerFlags = flags.Fcaller | flags.Ftimestamp | flags.FshortFile | flags.FLoggername
)

// BasicLogger implements Logger interface
type BasicLogger struct {
	//name logger name
	name string
	//filter logger filter, can be configured as composition of filters.
	filter filters.IFilter
	//appenders logger appenders, holds multiple appenders.
	appenders appenders.IAppenders
	//flags logger flags, each flag define logger feature.
	flags int
}

// NewConsoleLogger return a new Basic logger.
func NewConsoleLogger(name string, lvl level.Level, flags int) *BasicLogger {
	logger := BasicLogger{
		name: name,
		appenders: appenders.NewAppenders(
			appenders.NewConsoleAppender(
				encoders.NewRowEncoder(bufio.NewWriter(os.Stdout)))),
		filter: filters.NewLevelFilter(lvl),
		flags:  flags,
	}
	return &logger
}

// DefaultLogger return Basic logger with particular name.
func DefaultLogger(name string) *BasicLogger {
	level := level.DebugLevel //TODO:Should be taken from environment variable
	// TODO:Need function converting env-vars to Flags
	return NewConsoleLogger(name, level, FBasicLoggerFlags)
}

//getCaller return function caller, to gain max performance prefer to do not append caller to an event.
func getCaller(skip int, shortFileName bool) string {
	var b strings.Builder
	_, file, no, ok := runtime.Caller(skip)
	if ok {
		if shortFileName {
			if lastSlashIndex := strings.LastIndex(file, "/"); lastSlashIndex != -1 {
				file = file[lastSlashIndex+1:]
			}
		}
		fmt.Fprintf(&b, "%s:%d", file, no)
	}
	return b.String()
}

// enrichEvent enrich event with extended meta data.
func (l *BasicLogger) enrichEvent(e event.Event) {
	if e.ContainFlag(flags.Ftimestamp) {
		e.SetTimestamp(time.Now())
	}

	if e.ContainFlag(flags.Fcaller) {
		e.SetCaller(getCaller(4, e.ContainFlag(flags.FshortFile))) //Skip 4 depth levels, to get orig caller
	}

	if e.ContainFlag(flags.FLoggername) {
		e.SetLoggerName(l.name)
	}
}

// getName retrun logger name.
func (l *BasicLogger) getName() string {
	return l.name
}

// getFlags retrun logger flags.
func (l *BasicLogger) getFlags() int {
	return l.flags
}

// getAppenders retrun logger appenders.
func (l *BasicLogger) getAppenders() appenders.IAppenders {
	return l.appenders
}

//appendLogEvent Basci logger inner event appender mechnim.
func (l *BasicLogger) appendLogEvent(lvl level.Level, msg string) {
	e := event.NewBasicLogEvent(lvl, msg, l.flags)
	defer event.ReleaseLogEvent(e)

	if !l.filter.Enabled(e) {
		return
	}
	l.enrichEvent(e)
	l.appenders.DoAppendAll(e)
}

// Debug print message with Debug log level
func (l *BasicLogger) Debug(msg string) {
	l.appendLogEvent(level.DebugLevel, msg)
}

// Debugf format message fmt.Sprintf style, print to output with Debug level.
func (l *BasicLogger) Debugf(format string, vargs ...interface{}) {
	l.appendLogEvent(level.DebugLevel, fmt.Sprintf(format, vargs...))
}

// Debugln print message with Debug log level, append a new line.
func (l *BasicLogger) Debugln(msg string) {
	s := fmt.Sprintf(newLine, msg)
	l.appendLogEvent(level.DebugLevel, s)
}

// Info print message with Info log level
func (l *BasicLogger) Info(msg string) {
	l.appendLogEvent(level.InfoLevel, msg)
}

// Infof format message fmt.Sprintf style, print to output with Info level.
func (l *BasicLogger) Infof(format string, vargs ...interface{}) {
	l.appendLogEvent(level.InfoLevel, fmt.Sprintf(format, vargs...))
}

// Infoln print message with Info log level, append a new line.
func (l *BasicLogger) Infoln(msg string) {
	s := fmt.Sprintf(newLine, msg)
	l.appendLogEvent(level.InfoLevel, s)
}

// Warn print message with Warn log level
func (l *BasicLogger) Warn(msg string) {
	l.appendLogEvent(level.WarnLevel, msg)
}

// Warnf format message fmt.Sprintf style, print to output with Warn level.
func (l *BasicLogger) Warnf(format string, vargs ...interface{}) {
	l.appendLogEvent(level.WarnLevel, fmt.Sprintf(format, vargs...))
}

// Warnln print message with Warn log level, append a new line.
func (l *BasicLogger) Warnln(msg string) {
	s := fmt.Sprintf(newLine, msg)
	l.appendLogEvent(level.DebugLevel, s)
}

// Error print message with Error log level
func (l *BasicLogger) Error(msg string) {
	l.appendLogEvent(level.ErrorLevel, msg)
}

// Errorf format message fmt.Sprintf style, print to output with Error level.
func (l *BasicLogger) Errorf(format string, vargs ...interface{}) {
	l.appendLogEvent(level.ErrorLevel, fmt.Sprintf(format, vargs...))
}

// Errorln print message with Error log level, append a new line.
func (l *BasicLogger) Errorln(msg string) {
	s := fmt.Sprintf(newLine, msg)
	l.appendLogEvent(level.ErrorLevel, s)
}

// Panic print message with Panic log level, call panic
func (l *BasicLogger) Panic(msg string) {
	l.appendLogEvent(level.PanicLevel, msg)
	panic(msg)
}

// Panicf format message fmt.Sprintf style, print to output with Panic level, call panic.
func (l *BasicLogger) Panicf(format string, vargs ...interface{}) {
	s := fmt.Sprintf(format, vargs...)
	l.appendLogEvent(level.PanicLevel, s)
	panic(s)
}

// Panicln print message with Panic log level, append a new line, call panic.
func (l *BasicLogger) Panicln(msg string) {
	s := fmt.Sprintf(newLine, msg)
	l.appendLogEvent(level.PanicLevel, s)
	panic(s)
}

// Fatal print message with Fatal log level, followed by a call to os.Exit(1).
func (l *BasicLogger) Fatal(msg string) {
	l.appendLogEvent(level.FatalLevel, msg)
	os.Exit(1)
}

// Fatalf format message fmt.Sprintf style, print to output with Fatal level, followed by a call to os.Exit(1).
func (l *BasicLogger) Fatalf(format string, vargs ...interface{}) {
	l.appendLogEvent(level.FatalLevel, fmt.Sprintf(format, vargs...))
	os.Exit(1)
}

// Fatalln print message with Fatal log level, append a new line, followed by a call to os.Exit(1).
func (l *BasicLogger) Fatalln(msg string) {
	s := fmt.Sprintf(newLine, msg)
	l.appendLogEvent(level.FatalLevel, s)
	os.Exit(1)
}
