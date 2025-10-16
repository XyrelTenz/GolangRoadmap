package main

import (
	"encoding/json"
	"fmt"
)

// 🏠 Address represents a basic address structure.
type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

// 👤 User demonstrates a struct with a nested Address.
// - `omitempty` skips Age if it's 0.
// - `"-"` ignores Secret in JSON encoding.
type User struct {
	Name    string  `json:"full_name"`
	Age     int     `json:"age,omitempty"`
	Secret  string  `json:"-"`
	Address Address `json:"address"`
}

// 👤 FlatUser demonstrates an anonymous embedded struct.
// When marshaled to JSON, the fields of Address are merged (flattened).
type FlatUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Address     // 👈 anonymous embedding — no field name, so JSON flattens
}

func main() {
	// ===============================
	// ✅ Encode Struct with Nested JSON
	// ===============================
	u := User{
		Name:   "Xyrel",
		Age:    18,
		Secret: "hidden_password",
		Address: Address{
			City:    "Manila",
			Country: "Philippines",
		},
	}

	// MarshalIndent makes the output JSON pretty and readable.
	// Parameters:
	//   - u        → struct to encode
	//   - ""       → prefix (none)
	//   - "  "     → indent with 2 spaces
	jsonData, _ := json.MarshalIndent(u, "", "  ")

	fmt.Println("📌 Nested JSON Output:")
	fmt.Println(string(jsonData))
	/*
	   {
	     "full_name": "Xyrel",
	     "age": 18,
	     "address": {
	       "city": "Manila",
	       "country": "Philippines"
	     }
	   }
	*/

	// ===============================
	// 🔁 Decode JSON Back to Struct
	// ===============================
	var decodedUser User
	_ = json.Unmarshal(jsonData, &decodedUser)

	// After decoding, we can access struct fields normally.
	fmt.Println("📥 Decoded Struct:")
	fmt.Println("Name:", decodedUser.Name)
	fmt.Println("City:", decodedUser.Address.City)

	// ===============================
	// ✅ Encode Struct with Flattened JSON
	// ===============================
	fu := FlatUser{
		Name: "Xyrel",
		Age:  18,
		Address: Address{
			City:    "Manila",
			Country: "Philippines",
		},
	}

	flatJson, _ := json.MarshalIndent(fu, "", "  ")

	fmt.Println("\n📌 Flattened JSON Output:")
	fmt.Println(string(flatJson))
	/*
	   {
	     "name": "Xyrel",
	     "age": 18,
	     "city": "Manila",
	     "country": "Philippines"
	   }
	*/
}
