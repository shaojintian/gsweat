package worker

import (
	"github.com/shaojintian/gsweat/core"
	"time"
)
type Worker struct {

	// belong to what pool
	pool *core.Pool

	//binging func
	job chan func()

	//time when use and  reuse in the pool
	enterPool time.Time

}

type WorkerList []*Worker

func (w *Worker) run(f func()) {

}