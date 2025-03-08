package main

import ("testing"
"my-go-project/utils")

func TestGreet(t *testing.T) {
    expected := "Hello, Go!"
    result := utils.Greet("Go")
    if result != expected {
        t.Errorf("Expected '%s', but got '%s'", expected, result)
    }
}