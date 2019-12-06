package pool

//LRU FIFO

//generate new wk and dynamically adjust wksList size

type wkScheduler interface {
	isEmpty() bool
	len() int
	add(*Worker)
	optimize() int
}

type WorkerScheduler WorkerSched
type wksList []*Worker
type WorkerSched struct {
	wks wksList
}

func NewWorkerSched(size int) *WorkerSched {
	return &WorkerSched{
		wks: make([]*Worker, 0, size),
	}
}

func (wSched *WorkerScheduler) isEmpty() bool {
	if len(wSched.wks) == 0 {
		return true
	}
	return false
}
func (wSched *WorkerScheduler) len() int {
	return len(wSched.wks)
}
func (wSched *WorkerScheduler) add(wk *Worker) {
	wSched.wks = append(wSched.wks, wk)
}
func (wSched *WorkerScheduler) optimize() int {

	return wSched.len()
}
