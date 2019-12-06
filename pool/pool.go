package pool

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Pool struct {
	//size
	size int32

	//workers:all include working on and resting on
	wks WorkerList

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

func NewPool(size int32, extraFuncs ...extraFunc) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	// extra functions
	var extra *Extra
	for _, extraFunc := range extraFuncs {
		extraFunc(extra)
	}
	return &Pool{
		size:  size,
		close: PoolOpening,
		extra: extra,
	}, nil
}

// publish a func() to  do

func (p *Pool) PublishNewJob(f func()) error {
	if atomic.LoadUint32(&p.close) == PoolCLosed{
		return ErrPoolClosed
	}
	p.m.Lock()
	defer p.m.Unlock()

	if worker := p.wksPool.Get().(*Worker); worker != nil {
		ok := atomic.CompareAndSwapUint32(&worker.Status,Resting,Working)
		if !ok {return ErrConvertWokerStatus}
		//work
		worker.Job <- f
		worker.DoJob()
	}else{
		return errors.New("no valid worker")
	}
	return nil
}

func (p *Pool) ReUseWorker(worker *Worker) error{
	if atomic.LoadUint32(&p.close) == PoolCLosed{
		return ErrPoolClosed
	}
	//status change
	p.m.Lock()
	defer p.m.Unlock()
	if ok := atomic.CompareAndSwapUint32(&worker.Status,Working,Resting);!ok{
		return ErrConvertWokerStatus
	}
	p.wksPool.Put(worker)

	return nil

}

func (p *Pool) AddWkingNum (){
	atomic.AddInt64(&p.wkingNum, 1)
}

func (p *Pool) DecreWkingNum() {
	atomic.AddInt64(&p.wkingNum, -1)
}

func (p *Pool)GetWkingNum() int64 {
	return atomic.LoadInt64(&p.wkingNum)
}

func (p *Pool)Close() error {
	if ok := atomic.CompareAndSwapUint32(&p.close, PoolOpening, PoolCLosed);!ok{
		return errors.New("can not close pool")
	}

	return nil
}


