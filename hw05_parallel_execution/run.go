package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	tasksCh := make(chan Task)
	var errorCount int32
	for i := 0; i < n; i++ {
		wg.Add(1)
		go process(tasksCh, &errorCount, &wg)
	}
	checkErrorLimits := m > 0
	for _, task := range tasks {
		if checkErrorLimits && atomic.LoadInt32(&errorCount) >= int32(m) {
			break
		}
		tasksCh <- task
	}
	close(tasksCh)
	wg.Wait()
	if errorCount > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func process(tasksCh chan Task, errorCount *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasksCh {
		err := task()
		if err != nil {
			atomic.AddInt32(errorCount, 1)
		}
	}
}
