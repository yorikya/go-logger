/*
Package level represent logger event levels
*/
package level

//Level level struct
type Level int

const (
	//DebugLevel debug level instance
	DebugLevel Level = iota
	//InfoLevel info level instance
	InfoLevel
	//WarnLevel warn level instance
	WarnLevel
	//ErrorLevel error level instance
	ErrorLevel
	//PanicLevel panic level instance
	PanicLevel
	//FatalLevel fatal level instance
	FatalLevel
	//UndefinedLevel undefined level instance
	UndefinedLevel
)

//String string implementation
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	case UndefinedLevel:
		return "UNDEFINED"
	default:
		return "_LEVEL_"
	}
}
