package main

import (
	"fmt"
	"strings"
)

func main() {
    str := "hello, world"
    fmt.Println("Contains 'world':", strings.Contains(str, "world"))
    fmt.Println("To Upper:", strings.ToUpper(str))
    fmt.Println("Split:", strings.Split(str, ","))
}
