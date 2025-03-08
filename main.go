package main

import (
	"fmt"
	"net/http"
	"my-go-project/concurrency"
)

func main() {
	http.HandleFunc("/incrementWithRaceCondition", concurrency.IncrementHandlerForRaceCondition)

	http.HandleFunc("/incrementWithPessimisticLock", concurrency.IncrementHandlerForPessimisticLock)

	http.HandleFunc("/incrementWithOptimisticVersioning", concurrency.IncrementHandlerForOptimisticVersioning)

	http.HandleFunc("/valueWithRaceCondition", concurrency.ValueHandler)
	http.HandleFunc("/valueWithPessimisticLock", concurrency.ValueHandlerForPessimisticLock)
	http.HandleFunc("/valueWithOptimisticVersioning", concurrency.ValueHandlerForOptimisticVersioning)

	fmt.Println("Server is running on localhost:8080...")
	http.ListenAndServe("localhost:8080", nil)
}