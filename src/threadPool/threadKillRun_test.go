package threadPool

import (
	"github.com/remogatto/prettytest"
	"reflect"
)

type testSuiteThreadKill struct {
    prettytest.Suite
		testThread *threadKillRunner
}

func (t *testSuiteThreadKill) BeforeAll() {
	t.testThread = NewKillThread()
}

func (t *testSuiteThreadKill) TestThreadKillRun() {
	t.NotNil(t.testThread)
}

func (t *testSuiteThreadKill) TestThreadKillType() {
	test := reflect.TypeOf(t.testThread)
	sample := reflect.TypeOf(&threadKillRunner{})
  t.Equal(test, sample)
}

func (t *testSuiteThreadKill) TestThreadKillSig() {
	sig := t.testThread.Run()
  t.Equal(sig, KILL)
}
