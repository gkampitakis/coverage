package coverage

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"
)

type testingM interface {
	Run() int
}

var (
	// used for overriding during unit tests
	coverModeFunc = testing.CoverMode
	coverageFunc  = testing.Coverage
	writer        = io.Writer(os.Stdout)
	exitFunc      = os.Exit
)

func Run(t testingM, c float64, callbacks ...func()) {
	var code int
	defer func() {
		exitFunc(code)
	}()

	code = t.Run()

	if coverModeFunc() == "" {
		fmt.Fprintf(
			writer,
			"\nFAIL    coverage is not enabled. You can enable by using `-cover`\n\n",
		)
		code = 1
	} else {
		coverage := coverageFunc()
		if coverage*100 < c {
			code = 1
			fmt.Fprintf(writer, "\nFAIL    Coverage threshold not met %.1f >= %.1f for %s\n\n", c, coverage*100, packageName())
		}
	}

	for _, callback := range callbacks {
		callback()
	}
}

func packageName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	return funcName[:lastDot]
}
