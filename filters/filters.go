package filters

import (
	"gitlab.appsflyer.com/Architecture/af-logger-go/event"
	"gitlab.appsflyer.com/Architecture/af-logger-go/level"
)

type Filter interface {
	Enabled(event.Event) bool
}

type LevelFilter struct {
	filterLevel level.Level
}

func NewLevelFilter(lvl level.Level) LevelFilter {
	f := LevelFilter{
		filterLevel: lvl,
	}
	return f
}

func (l LevelFilter) Enabled(e event.Event) bool {
	return l.filterLevel >= e.GetLevel()
}
