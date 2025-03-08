package concurrency

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
	"time"
)

// Counter stores the current count and a version for optimistic locking.
type OptimisticCounter struct {
	value   int64
	version int64
}

// Global counter instance.
var optimisticCounter = &OptimisticCounter{value: 0, version: 1}

// Increment increases the counter value using optimistic locking.
func (c *OptimisticCounter) Increment() {
	for {
		// Read the current state atomically.
		currentValue := atomic.LoadInt64(&c.value)
		currentVersion := atomic.LoadInt64(&c.version)
		time.Sleep(1 * time.Millisecond)

		// Attempt to update the counter if the version hasn't changed.
		if atomic.CompareAndSwapInt64(&c.version, currentVersion, currentVersion+1) {
			atomic.StoreInt64(&c.value, currentValue+1)
			return
		}
	}
}

// Handler for the /increment endpoint.
func IncrementHandlerForOptimisticVersioning(w http.ResponseWriter, r *http.Request) {
	optimisticCounter.Increment()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"count":   atomic.LoadInt64(&optimisticCounter.value),
		"version": atomic.LoadInt64(&optimisticCounter.version),
	})
}

// Handler for the /value endpoint.
func ValueHandlerForOptimisticVersioning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"count":   atomic.LoadInt64(&optimisticCounter.value),
		"version": atomic.LoadInt64(&optimisticCounter.version),
	})
}
