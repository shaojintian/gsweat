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

func (ws *WorkerScheduler) isEmpty() bool {
	if len(ws.wks) == 0 {
		return true
	}
	return false
}
func (ws *WorkerScheduler) len() int {
	return len(ws.wks)
}
func (ws *WorkerScheduler) add(wk *Worker) {
	ws.wks = append(ws.wks, wk)
}
func (ws *WorkerScheduler) optimize() int {

	return ws.len()
}
