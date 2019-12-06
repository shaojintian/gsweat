package core

import (
	"errors"
	"math"
	"time"
)

// pool default config

const (
	DefaultPoolSize = math.MaxInt32

)


var (
	//errs
	ErrPoolClosed = errors.New("pool closed")
	ErrInvalidPoolSize = errors.New("pool size is invalid")

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

