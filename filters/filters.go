/*
Package filters contain interfaces and simple implementation for logger filters
Logger use filters to test if event should be append, basic filter is log level filter .
*/
package filters

import (
	"github.com/yorikya/go-logger/event"
	"github.com/yorikya/go-logger/level"
)

//IFilter filter interface
type IFilter interface {
	//Enabled return true when event can be append.
	Enabled(event.Event) bool
}

//LevelFilter implements filter interface
type LevelFilter struct {
	//filterLevel top filtered level
	filterLevel level.Level
}

// NewLevelFilter return a new  LevelFilter
// Args: filter top level
func NewLevelFilter(lvl level.Level) *LevelFilter {
	f := LevelFilter{
		filterLevel: lvl,
	}
	return &f
}

// Enabled interface function implementation
func (l LevelFilter) Enabled(e event.Event) bool {
	return l.filterLevel >= e.GetLevel()
}
