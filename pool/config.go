package pool

import (
	"errors"
	"time"
)

const (
	PoolCLosed uint32 = 0
	PoolOpening
	Resting
	Working
)


var (
	//errs
	ErrPoolClosed = errors.New("pool closed")
	ErrInvalidPoolSize = errors.New("pool size is invalid")
	ErrConvertWokerStatus = errors.New("convert woker status rest to worker error")

	//
)

//extra options
type extraFunc func(extra *Extra)

type Extra struct{

	expireTime time.Duration

}

func withNewExtra(e *Extra) extraFunc{
	return func(extra *Extra){
		extra = e
	}
}

func withExpireTime(t time.Duration) extraFunc{
	return func(extra *Extra){
		extra.expireTime = t
	}
}