package pool

import (
	"errors"
	"time"
)

const (
	PoolCLosed uint32 = iota
	PoolOpening
	Resting
	Working
)


var (
	//errs
	ErrPoolClosed = errors.New("goroutine pool closed")
	ErrInvalidPoolSize = errors.New("pool size is invalid")
	ErrConvertWorkerStatus = errors.New("convert woker status rest to worker error")

	//
)

//extra options
type extraFunc func(extra *Extra)

type Extra struct{

	expireTime time.Duration

}

func WithNewExtra(e *Extra) extraFunc{
	return func(extra *Extra){
		extra = e
	}
}

func WithExpireTime(t time.Duration) extraFunc{
	return func(extra *Extra){
		extra.expireTime = t
	}
}