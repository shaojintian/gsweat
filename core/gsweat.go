package core

import (
	"github.com/shaojintian/gsweat/pool"
	"time"
)

//init a default pool

var gsweat,_ =  pool.NewPool(DefaultPoolSize,pool.WithExpireTime(time.Duration(10*time.Second)))



