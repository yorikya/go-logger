/*
Package encoders contain interfaces and simple implementation for logger encoders
Logger use encoders to encode and stream data to particular stream out.
*/
package encoders

import (
	"bufio"
	"fmt"

	"github.com/yorikya/go-logger/event"
	"github.com/yorikya/go-logger/flags"
)

//IEncoder interface for encode an incoming event.
type IEncoder interface {
	// Encode an incoming event to string and stream to encoder out
	Encode(event.Event)
	//GetTimeFormat return as a string time format
	GetTimeFormat() string
	//GetWithLevel return true if encoder append level tag
	GetWithLevel() bool
}

//RowEncoder implements IEncoder interface
type RowEncoder struct {
	//out RowEncoder stream out
	out *bufio.Writer
	// timeFormat row time format
	timeFormat string
	// withLevel append to an incoming event level
	withLevel bool
}

//NewRowEncoder return an new RowEncoder
//Note: Get a bufio.Writer as an argument
func NewRowEncoder(w *bufio.Writer) *RowEncoder {
	e := RowEncoder{
		out:        w,
		timeFormat: "15:04:05.000",
		withLevel:  true,
	}
	return &e
}

//GetWithLevel return true if encoder configured to append log level.
func (enc *RowEncoder) GetWithLevel() bool {
	return enc.withLevel
}

// GetTimeFormat return time format as string
func (enc *RowEncoder) GetTimeFormat() string {
	return enc.timeFormat
}

//appendElement wrap a argument val with square parenthesis.
func (enc *RowEncoder) appendElement(val string) {
	enc.out.WriteByte('[')
	enc.appendElementVal(val)
	enc.out.WriteByte(']')
}

//appendElementVal writes a value given as an argument to encoder out
func (enc *RowEncoder) appendElementVal(val string) {
	fmt.Fprintf(enc.out, "%s", val)
}

//appendHeader append event header according to event flags.
func (enc *RowEncoder) appendHeader(evt event.Event) {
	if evt.ContainFlag(flags.Ftimestamp) {
		enc.appendElement(evt.GetTimestamp().Format(enc.timeFormat))
	}

	if enc.withLevel {
		enc.appendElement(evt.GetLevel().String())
	}

	if evt.ContainFlag(flags.Fcaller) {
		enc.appendElement(evt.GetCaller())
	}

	if evt.ContainFlag(flags.FLoggername) {
		enc.appendElement(evt.GetLoggerName())
	}

	enc.appendElementVal(evt.GetMessage())

}

//Encode encode event to string and stream to encoder out
func (enc *RowEncoder) Encode(evt event.Event) {
	enc.appendHeader(evt)

	enc.out.Flush()
}
