package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	work := NewPool(4)

	for i := 0; i < 50; i++ {
		work.Add(1)
		go testFunc(i, work)
	}

	fmt.Println("waiting...")
	work.Wait()
	t.Log("done")
}

func testFunc(i int, wg *WaitGroup) {
	defer wg.Done()
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"), "output: ", i)
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"), "output: ", i, "done")
}
