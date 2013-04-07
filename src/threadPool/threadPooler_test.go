package threadPool

import (
	"github.com/remogatto/prettytest"
	"reflect"
)

type testSuiteThreadPooler struct {
    prettytest.Suite
		testThread *ThreadPooler
		threadCount int
}

func (t *testSuiteThreadPooler) BeforeAll() {
	t.threadCount = 2
	t.testThread = NewThreadPooler(t.threadCount)
}

func (t *testSuiteThreadPooler) TestThreadPoolerNil() {
	t.NotNil(t.testThread)
}

func (t *testSuiteThreadPooler) TestThreadPoolerType() {
	test := reflect.TypeOf(t.testThread)
	control := reflect.TypeOf(&ThreadPooler{})
	t.Equal(test, control)
}

func (t *testSuiteThreadPooler) TestPoolSize() {
	t.Equal(t.threadCount, t.testThread.poolSize)
}

func (t *testSuiteThreadPooler) TestSetPoolSize() {
	t.testThread.SetPoolSize(t.threadCount + 1)
	t.NotEqual(t.threadCount, t.testThread.poolSize)
	t.testThread.SetPoolSize(t.threadCount)
}
