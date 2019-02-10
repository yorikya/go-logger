package impl

import (
	"bufio"
	"os"
	"runtime"
	"testing"
	"unicode"

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
	os.Stdout = writeFile

	go func() {
		scanner := bufio.NewScanner(readFile)
		for scanner.Scan() {
			// if txt := scanner.Text(); txt != "" {
			// 	c.out <- txt
			// }
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

func replaceDigit(source []rune, rep rune) string {
	var res []rune
	for _, r := range source {
		if unicode.IsDigit(r) {
			res = append(res, rep)
			continue
		}
		res = append(res, r)
	}
	return string(res)
}

func wrapElement(s string) string {
	return "[" + s + "]"
}

func assertEqual(t *testing.T, expect, current interface{}) {
	if expect != current {
		_, file, no, _ := runtime.Caller(2)
		t.Errorf("test failed expect: <%v>, current: <%v>\nCaller: %s, Line: %d", expect, current, file, no)
	}
}
func testOutput(t *testing.T, out string, lvl level.Level, logger *BasicLogger) {
	var seek, timstampLen int
	outRune := []rune(out)
	if flags.ContainFlag(logger.getFlags(), flags.Ftimestamp) {
		timeFmt := wrapElement(logger.getAppenders().GetAppender(0).GetEncoder().GetTimeFormat())
		timstampLen = len(timeFmt)
		seek = timstampLen
		assertEqual(t,
			replaceDigit([]rune(timeFmt), 'D'),
			replaceDigit(outRune[0:timstampLen], 'D'))
	}
	println(seek)
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
