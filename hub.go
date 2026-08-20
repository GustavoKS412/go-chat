package main

import "sync"

var (
	rooms = make(map[string]*room)
	mu    sync.Mutex
)

func getRoom(name string) *room {
	mu.Lock()
	defer mu.Unlock()

	if r, ok := rooms[name]; ok {
		return r
	}

	r := newRoom()
	rooms[name] = r
	go r.run()
	return r
}
