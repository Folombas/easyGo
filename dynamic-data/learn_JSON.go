package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	ID       int
	Username string
	phone    string
	Email    string
	birthday string
}

var jsonStr = `{"ID":1,"Username":"Bob","phone":"123456789","Email":"bob@example.com","birthday":"1970-01-01"}`

func main() {
	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("struct:\n\t%#v\n", u)

	u.phone = "987654321"
	result, _ := json.Marshal(u)
	fmt.Printf("json string:\n\t%s\n", string(result))
}