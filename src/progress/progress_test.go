package progress

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	p := NewProgress(0, 100)

	for i := 0; i <= 100; i++ {
		time.Sleep(10 * time.Millisecond)
		p.Add(int64(2))
	}
}

func TestDemo(t *testing.T) {
	fmt.Printf("\r123456")
	fmt.Printf("\rAB")
	fmt.Println()
}
