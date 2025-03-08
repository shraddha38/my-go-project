package concurrency;

import (
	"fmt"
	"net/http"
	"my-go-project/utils"
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
	utils.JsonResponseInt(w, map[string]int{
		"count": counter.value,
	})
}

func ValueHandler(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponseInt(w, map[string]int{
		"count": counter.value,
	})
	fmt.Printf("Current count: %d\n", counter.value)
}


