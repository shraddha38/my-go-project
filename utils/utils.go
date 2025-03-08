package utils

import (
    "encoding/json"
    "net/http"
)
func Greet(name string) string {
    return "Hello, " + name + "!"
}

func JsonResponse(w http.ResponseWriter, data interface{}) {

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(data)

}
func JsonResponseInt(w http.ResponseWriter, data interface{}) {

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(data)

}