package advanced

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var ErrInvalidWorkerCount = errors.New("worker count must be greater than zero")

type workItem struct {
	index int
	value int
}

type resultItem struct {
	index int
	value int
	err   error
}

// ParallelSquare computes square values concurrently while preserving output order.
func ParallelSquare(ctx context.Context, nums []int, workers int) ([]int, error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkerCount
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	jobs := make(chan workItem)
	results := make(chan resultItem)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					if job.value < 0 {
						results <- resultItem{index: job.index, err: fmt.Errorf("negative input at index %d", job.index)}
						continue
					}
					results <- resultItem{index: job.index, value: job.value * job.value}
				}
			}
		}()
	}

	go func() {
		defer close(jobs)
		for idx, n := range nums {
			select {
			case <-ctx.Done():
				return
			case jobs <- workItem{index: idx, value: n}:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	output := make([]int, len(nums))
	var firstErr error
	processed := 0

	for processed < len(nums) {
		select {
		case <-ctx.Done():
			if firstErr != nil {
				return nil, firstErr
			}
			return nil, ctx.Err()
		case res, ok := <-results:
			if !ok {
				if firstErr != nil {
					return nil, firstErr
				}
				return output, nil
			}
			processed++
			if res.err != nil && firstErr == nil {
				firstErr = res.err
				continue
			}
			output[res.index] = res.value
		}
	}

	if firstErr != nil {
		return nil, firstErr
	}
	return output, nil
}
