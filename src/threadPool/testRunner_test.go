package threadPool

import (
	"github.com/remogatto/prettytest"
	"testing"
)

func TestRunner(t *testing.T) {
	prettytest.RunWithFormatter(
			t,
			new(prettytest.TDDFormatter),
			new(testSuiteThreadKill),
			new(testSuite),
			new(testSuiteThreadPooler),
	)
}
