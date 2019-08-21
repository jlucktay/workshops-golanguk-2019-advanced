package main

import "sync"

// section: embed
type hits struct {
	sync.Mutex // embedding will promote all public methods
	n          int
}

func (h *hits) inc() {
	h.Lock()
	defer h.Unlock()
	h.n++
}

// section: embed
