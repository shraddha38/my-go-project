package concurrency

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
	"time"
	"fmt"
)

// Counter stores the current count and a version for optimistic Counter.
type OptimisticCounter struct {
	value   int64
	version int64
}

var optimisticCounter = &OptimisticCounter{value: 0, version: 1}

func (c *OptimisticCounter) Increment() {
	for {
		// Read the current state atomically.
		currentValue := atomic.LoadInt64(&c.value)
		currentVersion := atomic.LoadInt64(&c.version)

		// Sleep for a millisecond to avoid contention.
		time.Sleep(1 * time.Millisecond)

		// Attempt to update the counter if the version hasn't changed.
		if atomic.CompareAndSwapInt64(&c.version, currentVersion, currentVersion+1) {
			atomic.StoreInt64(&c.value, currentValue+1)
			return
		} else{
			fmt.Println("Optimistic increment failed")
			// If the version has changed, retry the operation.
		}
	}
}

func IncrementHandlerForOptimisticVersioning(w http.ResponseWriter, r *http.Request) {
	optimisticCounter.Increment()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"count":   atomic.LoadInt64(&optimisticCounter.value),
		"version": atomic.LoadInt64(&optimisticCounter.version),
	})
}

func ValueHandlerForOptimisticVersioning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"count":   atomic.LoadInt64(&optimisticCounter.value),
		"version": atomic.LoadInt64(&optimisticCounter.version),
	})
}
