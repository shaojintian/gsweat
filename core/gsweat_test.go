package core

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	n = 100000   // 10万
	MiB = 1048576
)
func TestGSweatPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++{
		_= gsweat.PublishNewJob(demoFunction)
	}
	wg.Wait()

	mem:= runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("gsweat pool used memory : %d MiB",mem.TotalAlloc/MiB)
	t.Logf("gsweat pool running workers: %d ",gsweat.GetWkingNum())
	//t.Logf("gsweat pool resting wks : %d MiB",mem.TotalAlloc/MiB)





}

func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++{
		go func() {
			demoFunction()
			wg.Done()
		}()
	}
	wg.Wait()

	mem:= runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	t.Logf("no pool used memory : %d MiB",mem.TotalAlloc/MiB)


}

func demoFunction(){
	time.Sleep(time.Duration(10) * time.Millisecond)
}