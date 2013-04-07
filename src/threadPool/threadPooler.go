package threadPool

func NewThreadPooler(numWorkers int) (tp *ThreadPooler) {
	tp = new(ThreadPooler)
	tp.SetPoolSize(numWorkers)
	tp.OpenChannel()
	return
}

type ThreadPooler struct {
	poolSize int
	channel  chan threadChannelRunner
}

func (s *ThreadPooler) PushToChannel(cm threadChannelRunner) {
	s.channel <- cm
}

func (s *ThreadPooler) SetPoolSize(i int) {
	s.poolSize = i
}

func (s *ThreadPooler) Join() {
	for i := 0; i < s.poolSize; i++ {
		s.PushToChannel(NewKillThread())
	}
}

func (s *ThreadPooler) OpenChannel() {
	s.channel = make(chan threadChannelRunner)
}

func (s *ThreadPooler) Start() {
	for i := 0; i < s.poolSize; i++ {

		go func() {

			for {
				channelMessage := <-s.channel

				if sig := channelMessage.Run(); sig == KILL {
					break
				}
			}
		}()
	}
}
