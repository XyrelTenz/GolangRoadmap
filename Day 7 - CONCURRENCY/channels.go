package main

import (
	"fmt"
	"time"
)

func main() {
    // Create a channel
    messages := make(chan string)

    // Start a goroutine that sends a message
    go func() {
        time.Sleep(1 * time.Second)
        messages <- "ping" // Send message to channel
    }()

    // Wait for the message from the channel
    msg := <-messages
    fmt.Println(msg)
}
