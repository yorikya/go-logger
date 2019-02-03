package encoders

import (
	"bufio"
	"fmt"

	"github.com/yorikya/go-logger/event"
	"github.com/yorikya/go-logger/flags"
)

const ()

type Encoder interface {
	Encode(event.Event)
}

type RowEncoder struct {
	out        *bufio.Writer
	timeFormat string
	withLevel  bool
}

func NewRowEncoder(w *bufio.Writer) *RowEncoder {
	e := RowEncoder{
		out:        w,
		timeFormat: "15:04:05.000",
		withLevel:  true,
	}
	return &e
}

func (enc *RowEncoder) appendElement(val string) {
	enc.out.WriteByte('[')
	enc.appendElementVal(val)
	enc.out.WriteByte(']')
}

func (enc *RowEncoder) appendElementVal(val string) {
	fmt.Fprintf(enc.out, "%s", val)
}

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

func (enc *RowEncoder) Encode(evt event.Event) {
	enc.appendHeader(evt)

	enc.out.WriteByte('\n')
	enc.out.Flush()

}
