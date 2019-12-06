package pool

import (
	"log"
	"time"
)
type Worker struct {

	// belong to what pool
	pool *Pool

	//binding func
	// we will reuse wk ,so use channel to dynamically binding func
	Job chan func()

	//time when use and  reuse in the pool
	reInPoolTime time.Time

	//status : work(1) or rest(0)
	Status uint32


}



func (w *Worker) DoJob() {
	go w.doJob()
}


func (w *Worker)doJob() {
	w.pool.AddWkingNum()

	if f,ok := <-w.Job;ok{
		f()
		w.pool.DecreWkingNum()
		err := w.pool.ReUseWorker(w)
		if err != nil {log.Print("reuse worker failed ", err)}
	}
}