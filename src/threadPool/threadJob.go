package threadPool

func NewThreadJob(r func()) (t *threadJob) {
	t = new(threadJob)
	t.SetRun = r
	return
}

type threadJob struct {
	SetRun func()
}

func (s *threadJob) Run() int {
	s.SetRun()
	return PASS
}
