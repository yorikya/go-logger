package encoders

import (
	"bufio"
	"fmt"

	"gitlab.appsflyer.com/Architecture/af-logger-go/event"
)

const ()

type Encoder interface {
	Encode(event.Event)
}

type RowEncoder struct {
	out        *bufio.Writer
	timeFormat string
}

func NewRowEncoder(w *bufio.Writer) *RowEncoder {
	e := RowEncoder{
		out:        w,
		timeFormat: "15:04:05.000",
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
	enc.appendElement(evt.GetTimestamp().Format(enc.timeFormat))
	enc.appendElement(evt.GetLevel().String())
	enc.appendElement(evt.GetCaller())
	enc.appendElement(evt.GetLoggerName())
	enc.appendElementVal(evt.GetMessage())
}

func (enc *RowEncoder) Encode(evt event.Event) {
	enc.appendHeader(evt)

	enc.out.WriteByte('\n')
	enc.out.Flush()

}
