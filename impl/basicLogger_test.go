package impl

import (
	"bufio"
	"os"
	"testing"

	"github.com/yorikya/go-logger/level"
)

type stdoutCapture struct {
	//out comunication out chan
	out chan string
	//readPipe read from file
	readPipe,
	//writePipe write to file
	writePipe,
	//origStdout holds original stdout file
	origStdout *os.File
}

func newStdoutCapture() *stdoutCapture {
	readFile, writeFile, err := os.Pipe()
	if err != nil {
		println("**** error: ", err.Error())
		return nil
	}

	c := &stdoutCapture{
		out:        make(chan string),
		readPipe:   readFile,
		writePipe:  writeFile,
		origStdout: os.Stdout,
	}

	go func() {
		scanner := bufio.NewScanner(readFile)
		for scanner.Scan() {
			c.out <- scanner.Text()
		}
	}()

	return c
}

func (c *stdoutCapture) getString() string {
	return <-c.out
}

func (c *stdoutCapture) close() {
	c.writePipe.Close()
	os.Stdout = c.origStdout
}

func testOutput(out string, lvl level.Level, flags int) {

}

func TestDebug(t *testing.T) {
	c := newStdoutCapture()
	defer c.close()

	l := NewConsoleLogger("Test", level.DebugLevel, FBasicLoggerFlags)
	l.Debug("test message")

	s := c.getString()

}
