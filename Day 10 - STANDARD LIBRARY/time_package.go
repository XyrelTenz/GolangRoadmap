package main

import (
	"fmt"
	"time"
)

func main() {
    now := time.Now()
    fmt.Println("Current time:", now)

    // Formatting time is done using a specific layout string.
    // The layout must be: Mon Jan 2 15:04:05 MST 2006
    fmt.Println("Formatted:", now.Format("2006-01-02 15:04:05"))
}
