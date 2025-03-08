package concurrency;

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Counter struct {
	value int
}

// Increment increases the counter value by 1.
func (c *Counter) Increment() {
	c.value++
}

var counter = &Counter{value: 0}


func IncrementHandlerForRaceCondition(w http.ResponseWriter, r *http.Request) {
	counter.Increment()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"count": counter.value,
	})
}

func ValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"count": counter.value,
	})
	fmt.Printf("Current count: %d\n", counter.value)
}


