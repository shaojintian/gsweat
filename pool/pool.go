package pool

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	//maximum number of goroutines
	size int64

	//workers:all include working on and resting on
	wks WorkerScheduler

	//to close(1) or not(0)
	close uint32

	//record working workers number
	wkingNum int64

	//sync
	//many sync.Mutex
	m sync.Mutex

	//many sync.Once
	once sync.Once

	//only sync.Cond
	cond *sync.Cond

	//[only one] sync.Pool to store resting workers
	wksPool *sync.Pool

	//extra functions
	extra *Extra
}

//constructor

func NewPool(size int64, extraFuncs ...extraFunc) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	// extra functions
	var extra *Extra
	for _, extraFunc := range extraFuncs {
		extraFunc(extra)
	}
	//init a pool
	p:= &Pool{
		size:  size,
		close: PoolOpening,
		extra: extra,
		wkingNum:0,
		m:sync.Mutex{},
	}
	p.wksPool = &sync.Pool{
		New: func() interface{} {
			return &Worker{
				pool:p,
				Job: nil,
				reInPoolTime:time.Now(),
				Status:Resting,
			}
		},
	}

	//async to dynamically adjust workers number
	//
	//go p.dynamicallyAdjustWorkers()
	//

	return p,nil
}

// publish a func() to  do

func (p *Pool) PublishNewJob(f func()) error {
	if atomic.LoadUint32(&p.close) == PoolCLosed {
		return ErrPoolClosed
	}

	log.Println("================startPublishNewJob================")
	// get Resting worker
	worker, err := p.GetRestingWorker()
	if worker != nil {
		log.Printf("[number %s]worker works ", worker.reInPoolTime)
		worker.Job <- f
		worker.DoJob()
	}

	return err
}

func (p *Pool) GetRestingWorker() (*Worker, error) {
	p.m.Lock()
	defer p.m.Unlock()
	if worker := p.wksPool.Get().(*Worker); worker != nil {
		return worker,nil
	}

	return nil, errors.New("no valid worker")

}

func (p *Pool) ReUseWorker(worker *Worker) error {
	if atomic.LoadUint32(&p.close) == PoolCLosed {
		return ErrPoolClosed
	}
	//status change
	p.m.Lock()
	defer p.m.Unlock()
	worker.reInPoolTime = time.Now()

	p.wksPool.Put(worker)

	return nil

}

func (p *Pool) AddWkingNum() {
	atomic.AddInt64(&p.wkingNum, 1)
}

func (p *Pool) DecreWkingNum() {
	atomic.AddInt64(&p.wkingNum, -1)
}

func (p *Pool) GetWkingNum() int64 {
	return atomic.LoadInt64(&p.wkingNum)
}

func (p *Pool) ClosePool() error {
	if ok := atomic.CompareAndSwapUint32(&p.close, PoolOpening, PoolCLosed); !ok {
		return errors.New("can not close pool")
	}

	return nil
}

func (p *Pool)dynamicallyAdjustWorkers(){

}