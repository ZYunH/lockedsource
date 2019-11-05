package lockedsource

import (
	"math/rand"
	"sync"
)

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source64
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

// New returns a new pseudo-random Source seeded with the given value.
// It is safe for concurrent use by multiple goroutines.
func New(seed int64) rand.Source {
	return rand.New(&lockedSource{src: rand.NewSource(seed).(rand.Source64)})
}
