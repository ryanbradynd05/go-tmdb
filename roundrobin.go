package tmdb

import (
	"sync"
)

// RoundRobin struct
type RoundRobin struct {
	currentTicker int
	maxAllowed    int
	mu            sync.Mutex
}

// InitRoundRobin func
func InitRoundRobin(maxAllowed int) RoundRobin {
	return RoundRobin{
		maxAllowed:    maxAllowed,
		currentTicker: 0,
		mu:            sync.Mutex{},
	}
}

// GetTicker func
func (r *RoundRobin) GetTicker() int {
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

// Next func
func (r *RoundRobin) Next() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentTicker < r.maxAllowed {
		r.currentTicker = r.currentTicker + 1
	} else {
		r.currentTicker = 0
	}
}
