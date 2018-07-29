package tmdb

import (
	"sync"
)

type roundRobin struct {
	currentTicker int
	maxAllowed int
	mu sync.Mutex
}

func InitRoundRobin(maxAllowed int) (roundRobin) {
	return roundRobin{
		maxAllowed: maxAllowed,
		currentTicker: 0,
		mu: sync.Mutex{},
	}
}

func (r *roundRobin) GetTicker() int  {
	r.mu.Lock()
	defer r.mu.Unlock()

	ticker := r.currentTicker

	if r.currentTicker < r.maxAllowed {
		r.currentTicker = r.currentTicker + 1
	} else {
		r.currentTicker = 0
	}

	return ticker
}

func (r *roundRobin) Next()  {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentTicker < r.maxAllowed {
		r.currentTicker = r.currentTicker + 1
	} else {
		r.currentTicker = 0
	}
}
