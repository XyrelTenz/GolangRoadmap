package main

import (
	"fmt"
	"sync"
)

func main() {
    var counter = 0
    var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock() // Lock the shared resource
            counter++
            mu.Unlock() // Unlock it
        }()
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("Final counter:", counter)
}
