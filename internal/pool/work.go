package pool

import (
	"log"
	job2 "scb-recipe-app/internal/job"
)

type Work struct {
	ID 	int
	Job string
}

type Worker struct {
	ID int
	WorkerChannel chan chan Work // used to communicate between dispatcher and workers
	Channel chan Work
	End chan bool
}

//start work
func (w *Worker) start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				job2.DoWork(job.Job, w.ID)
			case <-w.End:
				return
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}