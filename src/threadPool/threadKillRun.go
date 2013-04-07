package threadPool

func NewKillThread() *threadKillRunner {
	return new(threadKillRunner)
}

type threadKillRunner struct{}

func (s threadKillRunner) Run() int {
	return KILL
}
