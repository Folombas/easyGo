package main

import (
	"encoding/json"
	"fmt"
)

var jsonString = `[
	{"id":12, "username": "Gosha", "phone":25072025},
	{"id": "24", "address": "none", "company": "Best-code"}
]`

func main() {
	data := []byte(jsonString)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("unpacked in empty interface:\n%#v\n", user1)

	user2 := map[string]interface{}{
		"id":       12,
		"username": "Gosha",
	}
	var user2i interface{} = user2
	result, _ := json.Marshal(user2i)
	fmt.Printf("json string from map:\n %s\n", string(result))

}
