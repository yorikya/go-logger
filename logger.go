/*
Package logger, interface for app logs.
*/
package logger //v1

// Logger basic logger API
type Logger interface {
	// Debug string message marked with debug level
	Debug(string)
	// Debugf substitute vargs to string format fmt.Sprintf style with debug level
	Debugf(string, ...interface{})
	// Debugln string message with debug level, append a new line symbol.
	Debugln(string)
	//Info string message marked with info level
	Info(string)
	// Infof substitute vargs to string format fmt.Sprintf style with info level
	Infof(string, ...interface{})
	// Infoln string message with info level, append a new line symbol.
	Infoln(string)
	// Warn string message marked with warn level
	Warn(string)
	// Warnf substitute vargs to string format fmt.Sprintf style with warn level
	Warnf(string, ...interface{})
	// Warnln string message with warn level, append a new line symbol.
	Warnln(string)
	// Error string message marked with error level
	Error(string)
	// Errorf substitute vargs to string format fmt.Sprintf style with error level
	Errorf(string, ...interface{})
	// Errorln string message with error level, append a new line symbol.
	Errorln(string)
	// Panic string message marked with panic level, call panic()
	Panic(string)
	// Panicf substitute vargs to string format fmt.Sprintf style with panic level, call panic()
	Panicf(string, ...interface{})
	// Panicln string message with panic level, append a new line symbol, call panic()
	Panicln(string)
	// Fatal string message marked with fatal level, call os.exit(1)
	Fatal(string)
	// Fatalf substitute vargs to string format fmt.Sprintf style with fatal level, call os.exit(1)
	Fatalf(string, ...interface{})
	// Fatalln string message with fatal level, append a new line symbol, call os.exit(1)
	Fatalln(string)
}

// LoggerFactory logger factory API
type LoggerFactory interface {
	GetLogger(string) Logger
}
