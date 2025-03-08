package concurrency

import (
	"net/http"
	"sync"
	"my-go-project/utils"
)

// Counter stores the current count and a mutex for thread safety.
type PessimisticCounter struct {
	value int
	mutex sync.Mutex
}

// Increment increases the counter value by 1 using pessimistic locking.
func (c *PessimisticCounter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

// GetValue returns the current counter value.
func (c *PessimisticCounter) GetPessimisticValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

var pessimisticCounter = &PessimisticCounter{value: 0}

func IncrementHandlerForPessimisticLock(w http.ResponseWriter, r *http.Request) {
	pessimisticCounter.Increment()
	utils.JsonResponseInt(w, map[string]int{
		"count": pessimisticCounter.GetPessimisticValue(),
	})
}

func ValueHandlerForPessimisticLock(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponseInt(w, map[string]int{
		"count": pessimisticCounter.GetPessimisticValue(),
	})
}
