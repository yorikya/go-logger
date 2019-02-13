package impl

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"unicode"

	"github.com/yorikya/go-logger/encoders"
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

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
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
		scanner.Split(ScanCRLF)

		for scanner.Scan() {
			msg := scanner.Text()

			c.out <- msg
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

func replaceDigitWithD(source []rune) string {
	return replaceDigit(source, 'D')
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

func assertTrue(t *testing.T, current interface{}) {
	assertEqual(t, true, current)
}

func cutFirstElement(s string) string {
	start := strings.IndexRune(s, '[')
	end := strings.IndexRune(s, ']') + 1
	return string([]rune(s)[start:end])
}

func testOutput(t *testing.T, msg, out string, lvl level.Level, logger *BasicLogger) {
	var seek, timstampLen int
	outRune := []rune(out)
	//Test timestamp
	if flags.ContainFlag(logger.getFlags(), flags.Ftimestamp) {
		timeFmt := wrapElement(logger.getAppenders().GetAppender(0).GetEncoder().GetTimeFormat())
		timstampLen = len(timeFmt)

		assertEqual(t,
			replaceDigitWithD([]rune(timeFmt)),
			replaceDigitWithD(outRune[seek:timstampLen]))
		seek += timstampLen
	}

	//Test log level
	if logger.getAppenders().GetAppender(0).GetEncoder().GetWithLevel() {
		levelFmt := wrapElement(lvl.String())
		levelLen := len(levelFmt)

		assertEqual(t, levelFmt, string(outRune[seek:seek+levelLen]))
		seek += levelLen
	}

	//Test caller
	if flags.ContainFlag(logger.getFlags(), flags.Fcaller) {
		caller := cutFirstElement(string(outRune[seek:]))
		callerLen := len(caller)

		assertTrue(t, strings.Contains(caller, ".go"))

		if !flags.ContainFlag(logger.getFlags(), flags.FshortFile) {
			assertTrue(t, strings.Contains(caller, "/"))
		}

		seek += callerLen
	}

	//Test logger name
	if flags.ContainFlag(logger.getFlags(), flags.FLoggername) {
		name := wrapElement(logger.getName())
		nameLen := len(name)
		assertEqual(t, name, string(outRune[seek:seek+nameLen]))

		seek += nameLen
	}

	//Test message
	assertEqual(t, msg, string(outRune[seek:]))
}

func TestDebug(t *testing.T) {
	c := newStdoutCapture()
	flags := FBasicLoggerFlags
	msg := "test message"

	l := NewConsoleLogger("Test", level.DebugLevel, flags)
	l.Debug(msg)

	c.close()
	testOutput(t, msg, c.getString(), level.DebugLevel, l)
}

func TestDebugf(t *testing.T) {
	c := newStdoutCapture()
	flags := FBasicLoggerFlags
	msg := "test %s with %d substitutions"
	firstArg := "mesage"
	secondArg := 2

	l := NewConsoleLogger("Test", level.DebugLevel, flags)
	l.Debugf(msg, firstArg, secondArg)

	c.close()
	testOutput(t, fmt.Sprintf(msg, firstArg, secondArg), c.getString(), level.DebugLevel, l)
}

func TestDebugln(t *testing.T) {
	c := newStdoutCapture()
	flags := FBasicLoggerFlags
	msg := "test message with new line"

	l := NewConsoleLogger("Test", level.DebugLevel, flags)
	l.Debugln(msg)

	c.close()
	testOutput(t, msg+encoders.NewLine, c.getString(), level.DebugLevel, l)
}
