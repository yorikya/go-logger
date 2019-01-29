/*
*** Appsflyer 2018 ***
***  Platform Team   ***
***  APIGateway    ***
Package logger, interface for app logs.
*/
package logger //v1

//Logger API
type Logger interface {
	//Debug string message marked with debug level
	Debug(string)
	// //Debugf substitute vargs to string format fmt.Sprintf style with debug level
	// Debugf(string, ...interface{})
	// //DebugJSON string message with debug level and expands JSON context.
	// DebugJSON(string, map[string]interface{})
	// //Info string message marked with info level
	// Info(string)
	// //Infof substitute vargs to string format fmt.Sprintf style with info level
	// Infof(string, ...interface{})
	// //InfoJSON string message with info level and expands JSON context.
	// InfoJSON(string, map[string]interface{})
	// //Warn string message marked with warn level
	// Warn(string)
	// //Warnf substitute vargs to string format fmt.Sprintf style with warn level
	// Warnf(string, ...interface{})
	// //WarnJSON string message with warn level and expands JSON context.
	// WarnJSON(string, map[string]interface{})
	// //Error string message marked with error level
	// Error(string)
	// //Errorf substitute vargs to string format fmt.Sprintf style with error level
	// Errorf(string, ...interface{})
	// //ErrorJSON string message with error level and expands JSON context.
	// ErrorJSON(string, map[string]interface{})
	// //Panic string message marked with panic level, call panic()
	// Panic(string)
	// //Panicf substitute vargs to string format fmt.Sprintf style with panic level, call panic()
	// Panicf(string, ...interface{})
	// //PanicJSON string message with panic level and expands JSON context, call panic()
	// PanicJSON(string, map[string]interface{})
	// //Fatal string message marked with fatal level, call os.exit(1)
	// Fatal(string)
	// //Fatalf substitute vargs to string format fmt.Sprintf style with fatal level, call os.exit(1)
	// Fatalf(string, ...interface{})
	// //FatalJSON string message with fatal level and expands JSON context, call os.exit(1)
	// FatalJSON(string, map[string]interface{})
}

type LoggerFactory interface {
	GetLogger(string) Logger
}
