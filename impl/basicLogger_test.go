package impl

import (
	"bufio"
	"os"
	"testing"

	"github.com/yorikya/go-logger/flags"
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
	close(c.out)
	os.Stdout = c.origStdout
}

func testOutput(t *testing.T, out string, lvl level.Level, logger *BasicLogger) {
	if flags.ContainFlag(logger.getFlags(), flags.Ftimestamp) {
		timeFmt := logger.getAppenders().GetAppender(0).GetEncoder().GetTimeFormat()
		println("the time format:", timeFmt)
		// enc.appendElement(evt.GetTimestamp().Format(enc.timeFormat))
	}

	// if withLevel {
	// 	// enc.appendElement(evt.GetLevel().String())
	// }

	// if flags.ContainFlag(loggerFlags, flags.Fcaller) {
	// 	getCaller(4, flags.ContainFlag(loggerFlags, flags.FshortFile))
	// 	// enc.appendElement(evt.GetCaller())
	// }

	// if flags.ContainFlag(loggerFlags, flags.FLoggername) {
	// 	// enc.appendElement(evt.GetLoggerName())
	// }

	// //Test message
	// // enc.appendElementVal(evt.GetMessage())
}

func TestDebug(t *testing.T) {
	c := newStdoutCapture()
	flags := FBasicLoggerFlags
	// defer c.close()

	l := NewConsoleLogger("Test", level.DebugLevel, flags)
	l.Debug("test message")
	c.close()
	testOutput(t, c.getString(), level.DebugLevel, l)
}
