package sync

import "sync/atomic"

var total int32 = 1

func add(i int32) {
	atomic.AddInt32(&total, i)
}

func get() int32 {
	return total
}
