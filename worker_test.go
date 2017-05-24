package worker_test

import (
	"sync/atomic"
	"testing"

	"github.com/korovkin/worker"
)

func TestWorker(t *testing.T) {
	C := int64(0)

	w := worker.NewWorkerDefault("test")
	w.Enqueue(func() {
		atomic.AddInt64(&C, 1000000)
		t.Log(" C += ", 1000000)
	})

	const N = 10

	for i := 0; i < N; i++ {
		w.Enqueue(func() {
			t.Log(" C += ", 1, i)
			atomic.AddInt64(&C, 1)
		})
	}

	w.EnqueueSync(func() {
		atomic.AddInt64(&C, 1000)
		t.Log(" C += ", 1000)
	})

	if C != 1000000+1000+N {
		t.Fatal("failed test with C:", C)
		t.FailNow()
	}

	t.Log("PASS: C:", C)
}
