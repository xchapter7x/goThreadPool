package main

import (
	"../threadPool"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runThreads()
}

func runThreads() {
	numWorkers := runtime.NumCPU() * 10
	runtime.GOMAXPROCS(numWorkers)
	tp := threadPool.NewThreadPooler(numWorkers)
	tp.Start()
	arr := []int{2,6,1,7,4,2,1,4,2,2,2,4,2}

	for _, val := range arr {
		x := val
		f := func() {
			func(i int) {
				time.Sleep(time.Duration(i) * time.Second)
				fmt.Println(i)
			}(x)
		}
		job := threadPool.NewThreadJob(f)
		tp.PushToChannel(job)
	}
	tp.Join()
}
