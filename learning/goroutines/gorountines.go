package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main() {
    // wg.Add(2) // Number of Goroutines to wait for
    go worker(1)
    go worker(2)
    // wg.Wait() // Wait for all Goroutines to finish
    fmt.Println("FInished")
}

func worker(id int) {
    fmt.Println("Before", id)
    // defer wg.Done()
    fmt.Println("after", id)
}
