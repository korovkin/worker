package worker

import (
	"log"
)

type WorkerJob struct {
	Job func()

	stop        bool
	syncChannel chan int
}

type Worker struct {
	Name string

	queue chan WorkerJob
}

func NewWorkerDefault(name string) *Worker {
	return NewWorker(name, make(chan WorkerJob, 1000))
}

func NewWorker(name string, queue chan WorkerJob) *Worker {
	if queue == nil {
		queue = make(chan WorkerJob, 1000)
	}

	workerLoop := func(w Worker) {
		for {
			j := <-w.queue
			j.Job()
			if j.stop {
				break
			}

			if j.syncChannel != nil {
				j.syncChannel <- 0xC0FFE
			}
		}
	}

	w := Worker{Name: name, queue: queue}
	go workerLoop(w)
	return &w
}

func (w *Worker) Close() {
	w.Stop(func() {})
}

func (w *Worker) Stop(job func()) {
	j := WorkerJob{Job: job, stop: true, syncChannel: nil}
	w.queue <- j
}

func (w *Worker) Enqueue(job func()) {
	j := WorkerJob{Job: job, stop: false, syncChannel: nil}
	w.queue <- j
}

func (w *Worker) EnqueueSync(job func()) {
	j := WorkerJob{Job: job, stop: false, syncChannel: make(chan int)}
	w.queue <- j
	coffee := <-j.syncChannel

	if coffee != 0xC0FFE {
		log.Fatalln(coffee)
	}
}

func (w *Worker) GetQueueSize() int {
	return len(w.queue)
}
