package coverage

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

func Run(t *testing.M, c float64, callbacks ...func(t *testing.M)) {
	var code int
	defer func() {
		os.Exit(code)
	}()

	code = t.Run()

	if testing.CoverMode() == "" {
		fmt.Printf("\nFAIL    coverage is not enabled. You can enable by using `-cover`\n\n")
		code = 1
	} else {
		coverage := testing.Coverage()
		if coverage*100 < c {
			code = 1
			fmt.Printf("\nFAIL    Coverage threshold not met %.1f >= %.1f for %s\n\n", c, coverage*100, packageName())
		}
	}

	for _, callback := range callbacks {
		callback(t)
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
