package coverage

import (
	"strings"
	"testing"
)

func mockExit(t *testing.T, code int) func(int) {
	t.Helper()

	return func(c int) {
		if code != c {
			t.Errorf("[expected]: %d \n[received]:%d\n", code, c)
		}
	}
}

type mockTestingM struct {
	code int
}

func (m mockTestingM) Run() int {
	return m.code
}

func TestCoverage(t *testing.T) {
	t.Run("should run tests successfully", func(t *testing.T) {
		exitFunc = mockExit(t, 0)
		coverModeFunc = func() string { return "mock-mode" }
		coverageFunc = func() float64 { return 1 }
		m := mockTestingM{code: 0}

		Run(m, 100)
	})

	t.Run("should fail if no cover enabled", func(t *testing.T) {
		exitFunc = mockExit(t, 1)
		coverModeFunc = func() string { return "" }
		coverageFunc = func() float64 { return 1 }
		m := mockTestingM{code: 0}
		s := strings.Builder{}
		writer = &s

		Run(m, 100)
		if s.String() != "\nFAIL    coverage is not enabled. You can enable by using `-cover`\n\n" {
			t.Errorf("expected coverage not enabled msg but got: %s\f", s.String())
		}
	})

	t.Run("should fail if coverage is not met", func(t *testing.T) {
		exitFunc = mockExit(t, 1)
		coverModeFunc = func() string { return "mock-mode" }
		coverageFunc = func() float64 { return 0.5 }
		m := mockTestingM{code: 0}
		s := strings.Builder{}
		writer = &s

		Run(m, 51)
		if s.String() != "\nFAIL    Coverage threshold not met 51.0 >= 50.0 for coverage.TestCoverage\n\n" {
			t.Errorf("expected coverage threshold error but got: %s\n", s.String())
		}
	})

	t.Run("should call callback funcs", func(t *testing.T) {
		exitFunc = mockExit(t, 0)
		coverModeFunc = func() string { return "mock-mode" }
		coverageFunc = func() float64 { return 1 }
		m := mockTestingM{code: 0}
		called := 0

		Run(m, 100, func() {
			called++
		})

		if called != 1 {
			t.Errorf("callback expected to be called once but called: %d\n", called)
		}
	})
}
