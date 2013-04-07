package threadPool

import (
	"github.com/remogatto/prettytest"
	"reflect"
)

type testSuite struct {
    prettytest.Suite
		testThread *threadJob
}

func (t *testSuite) After() {

}

func (t *testSuite) BeforeAll() {
	f := func() {
		//test sample func	
	}
	t.testThread = NewThreadJob(f)
}

func (t *testSuite) TestThreadJobNil() {
	t.NotNil(t.testThread)
}

func (t *testSuite) TestThreadJobReturnType() {
	test := reflect.TypeOf(t.testThread)
	sample := reflect.TypeOf(&threadJob{})
	t.Equal(test, sample)
}

func (t *testSuite) TestThreadJobSetRunNil() {
	t.NotNil(t.testThread.SetRun)
}
