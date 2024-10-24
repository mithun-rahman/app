package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

func NewPool(workerCount int) *Pool {
	pool := &Pool{
		workQueue: make(chan Job, workerCount),
	}
	pool.wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer pool.wg.Done()
			for job := range pool.workQueue {
				job()
			}
		}()
	}
	return pool
}

func (p *Pool) AddJob(job Job) {
	p.workQueue <- job
}

func (p *Pool) Wait() {
	close(p.workQueue)
	p.wg.Wait()
}

func ThreadPool() {
	pool := NewPool(3)

	for i := 0; i < 10; i++ {
		job := func() {
			time.Sleep(2 * time.Second)
			fmt.Println(fmt.Sprintf("thread %d", i))
		}
		pool.AddJob(job)
	}
	pool.Wait()
}
