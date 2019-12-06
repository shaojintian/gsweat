package core

import (
	wk "github.com/shaojintian/gsweat/worker"
	"sync"
)

type Pool struct {
	//size
	size int32

	//workers
	wks wk.WorkerList

	//to close or not
	close bool

	//sync
	//many sync.Mutex
	m sync.Mutex

	//many sync.Once
	once sync.Once

	//only sync.Cond
	cond *sync.Cond

	//only sync.Pool to store workers
	wksPool *sync.Pool

	//extra functions
	extra *Extra

}

//constructor

func NewPool (size int32,extraFuncs ...extraFunc ) (*Pool,error) {
	if size <=0 {
		return nil,ErrInvalidPoolSize
	}
	// extra functions
	var extra *Extra
	for _, extraFunc := range extraFuncs{
		extraFunc(extra)
	}
	return &Pool{
		size:size,
		close:false,
		extra:extra,
	},nil
}

// publish a func() to  do

func (p *Pool) publish(f func()) error{

}

func(p *Pool) getWorker() error{

}