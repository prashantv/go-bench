package main

import "sync"

// Test for overhead and alternatives to defer.
// See http://lk4d4.darth.io/posts/defer/

// SyncVal is a value protected by an interface.
type SyncVal struct {
	sync.Mutex
	v interface{}
}

// PutSimple gets a lock and unlocks it explicitly after the operation.
func (sv *SyncVal) PutSimple(v interface{}) {
	sv.Lock()
	sv.v = v
	sv.Unlock()
}

// PutDefer uses a defer to lock and unlock the mutex.
func (sv *SyncVal) PutDefer(v interface{}) {
	sv.Lock()
	defer sv.Unlock()

	sv.v = v
}

func (sv *SyncVal) withMutex(f func()) {
	sv.Lock()
	f()
	sv.Unlock()
}

// PutFunc uses a function to lock and unlock.
func (sv *SyncVal) PutFunc(v interface{}) {
	sv.withMutex(func() {
		sv.v = v
	})
}
