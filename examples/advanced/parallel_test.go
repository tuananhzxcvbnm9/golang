package advanced

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestParallelSquareSuccess(t *testing.T) {
	ctx := context.Background()
	nums := []int{1, 2, 3, 4}

	got, err := ParallelSquare(ctx, nums, 2)
	if err != nil {
		t.Fatalf("ParallelSquare() error = %v", err)
	}

	want := []int{1, 4, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParallelSquare() = %v, want %v", got, want)
	}
}

func TestParallelSquareInvalidWorkers(t *testing.T) {
	_, err := ParallelSquare(context.Background(), []int{1, 2}, 0)
	if !errors.Is(err, ErrInvalidWorkerCount) {
		t.Fatalf("expected ErrInvalidWorkerCount, got %v", err)
	}
}

func TestParallelSquareCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := ParallelSquare(ctx, []int{1, 2, 3}, 2)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}

func TestParallelSquareNegativeInput(t *testing.T) {
	_, err := ParallelSquare(context.Background(), []int{2, -1, 3}, 2)
	if err == nil {
		t.Fatal("expected error for negative input")
	}
}

func BenchmarkParallelSquareWorkerCounts(b *testing.B) {
	nums := make([]int, 10_000)
	for i := range nums {
		nums[i] = i + 1
	}

	workerCounts := []int{1, 2, 4, 8}
	for _, workers := range workerCounts {
		workers := workers
		b.Run("workers_"+strconv.Itoa(workers), func(b *testing.B) {
			ctx := context.Background()
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := ParallelSquare(ctx, nums, workers)
				if err != nil {
					b.Fatalf("ParallelSquare() error = %v", err)
				}
			}
		})
	}
}
