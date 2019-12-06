package core

import "github.com/shaojintian/ghurri-net/worker"

type Pool struct  {
	//size
	size int32

	//workers
	wks []*worker


}