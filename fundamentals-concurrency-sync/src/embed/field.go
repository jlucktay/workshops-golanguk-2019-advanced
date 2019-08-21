package main

import "sync"

// section: embed
type hits struct {
	mu sync.Mutex // methods are no longer promoted and only accessible within this package now.
	n  int
}

func (h *hits) inc() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
}

// section: embed
