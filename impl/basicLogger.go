//Should be moved to independnt module.
package impl

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"gitlab.appsflyer.com/Architecture/af-logger-go/appenders"
	"gitlab.appsflyer.com/Architecture/af-logger-go/encoders"
	"gitlab.appsflyer.com/Architecture/af-logger-go/event"
	"gitlab.appsflyer.com/Architecture/af-logger-go/filters"
	"gitlab.appsflyer.com/Architecture/af-logger-go/level"
	"gitlab.appsflyer.com/Architecture/af-logger-go/mdc"
)

const (
	Fcaller = 1 << iota
	Ftimestamp
	Fshortfile
	FLoggername

	FBasicLoggerFlags = Fcaller | Ftimestamp | Fshortfile | FLoggername
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

func (l *BasicLogger) containFlag(featureFlag int) bool {
	return l.flags&Ftimestamp != 0
}

func (l *BasicLogger) enrichEvent(e event.Event) {
	switch {
	case l.containFlag(Ftimestamp):
		e.SetTimestamp(time.Now())
		fallthrough

	case l.containFlag(Fcaller):
		e.SetCaller(getCaller(4, l.containFlag(Fshortfile))) //Skip 4 depth levels, to get orig caller
		fallthrough

	case l.containFlag(FLoggername):
		e.SetLoggerName(l.name)
	}
}

func (l *BasicLogger) appendLogEvent(lvl level.Level, msg string) {
	e := event.NewBasicLogEvent(lvl, msg)
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
