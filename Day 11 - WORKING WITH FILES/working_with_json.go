package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // Marshalling (Go struct to JSON)
    user1 := User{Name: "Alice", Age: 30}
    jsonData, err := json.Marshal(user1)
    if err != nil {
        panic(err)
    }
    fmt.Println("JSON:", string(jsonData))

    // Unmarshalling (JSON to Go struct)
    var user2 User
    jsonStr := `{"name":"Bob","age":25}`
    err = json.Unmarshal([]byte(jsonStr), &user2)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Unmarshalled user: %+v\n", user2)
}
