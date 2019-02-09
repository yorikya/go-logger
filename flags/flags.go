/*
Package flags define global flags for Logger/Event
Logger create an new event from each incoming message, represnet event behavior .
*/
package flags

const (
	//Fcaller caller flag. Note: To gain best performance disable caller.
	Fcaller = 1 << iota
	//Ftimestamp timestamp flag. Append timestamp to event.
	Ftimestamp
	//FshortFile short file flag. Cut absolute file path to file name only.
	FshortFile
	//FLoggername logger nme flag. Append logger name to an event.
	FLoggername
)
