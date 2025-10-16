package main

import (
	"fmt"
	"time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    // Run this function concurrently
    go say("world")

    say("hello")
}
