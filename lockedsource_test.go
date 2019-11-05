package lockedsource

import (
	"sync"
	"testing"
	"time"
)

func TestLockedSource(t *testing.T) {
	var wg sync.WaitGroup
	source := New(time.Now().UnixNano())
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			source.Seed(source.Int63())
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
