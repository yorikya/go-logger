//Should be moved to independnt module.
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
	"github.com/yorikya/go-logger/mdc"
)

const (
	FBasicLoggerFlags = flags.Fcaller | flags.Ftimestamp | flags.Fshortfile | flags.FLoggername
)

type BasicLogger struct {
	name      string
	filter    filters.Filter
	appenders appenders.IAppenders
	flags     int
	mdc       mdc.IMDC
}

func ConsoleLogger(name string, lvl level.Level, flags int, mdc mdc.IMDC) *BasicLogger {
	logger := BasicLogger{
		name: name,
		appenders: appenders.NewAppenders(
			appenders.NewConsoleAppender(
				encoders.NewRowEncoder(bufio.NewWriter(os.Stdout)))),
		filter: filters.NewLevelFilter(lvl),
		flags:  flags,
		mdc:    mdc,
	}
	return &logger
}

func DefaultLogger(name string) *BasicLogger {
	level := level.DebugLevel //TODO:Should be taken from environment variable
	// TODO:Need function converting env-vars to Flags
	mdc := mdc.NewContext()
	return ConsoleLogger(name, level, FBasicLoggerFlags, mdc)
}

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

func (l *BasicLogger) enrichEvent(e event.Event) {
	if e.ContainFlag(flags.Ftimestamp) {
		e.SetTimestamp(time.Now())
	}

	if e.ContainFlag(flags.Fcaller) {
		e.SetCaller(getCaller(4, e.ContainFlag(flags.Fshortfile))) //Skip 4 depth levels, to get orig caller
	}

	if e.ContainFlag(flags.FLoggername) {
		e.SetLoggerName(l.name)
	}
}

func (l *BasicLogger) appendLogEvent(lvl level.Level, msg string) {
	e := event.NewBasicLogEvent(lvl, msg, l.flags)
	defer event.ReleaseLogEvent(e)

	if !l.filter.Enabled(e) {
		return
	}
	l.enrichEvent(e)
	l.appenders.DoAppendAll(e)
}

func (l *BasicLogger) Debug(msg string) {
	l.appendLogEvent(level.DebugLevel, msg)
}
