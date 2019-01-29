/*
*** Appsflyer 2018 ***
***  Platform Team   ***
***  Logger Unit Test    ***
 */
package logger

import (
	"runtime"
	"testing"
)

const (
	debugLevel = "DEBUG"
	infoLevel  = "INFO"
	warnLevel  = "WARN"
	errorLevel = "ERROR"
)

func assertEqual(t *testing.T, expect, current interface{}) {
	if expect != current {
		_, file, no, _ := runtime.Caller(2)
		t.Errorf("test failed expect: <%v>, current: <%v>\nCaller: %s, Line: %d", expect, current, file, no)
	}
}

func printTest(log Logger) {
	log.Debug("printTest function!!!!")
}

func TestDefaultInit(t *testing.T) {
	log := GetLogger("ROOT")

	log.Debug("Test Message")

	for i := 1; i <= 100; i++ {
		go printTest(log)
	}

}
