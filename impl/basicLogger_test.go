package impl

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/yorikya/go-logger/level"
)

func getStdout() (chan string, *os.File) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	return outC, w
}
func TestDebugFamily(t *testing.T) {
	old := os.Stdout
	o, w := getStdout()
	l := NewConsoleLogger("Test", level.DebugLevel, FBasicLoggerFlags)
	l.Debug("test message")

	w.Close()
	os.Stdout = old
	println("This is my output", <-o)

}
