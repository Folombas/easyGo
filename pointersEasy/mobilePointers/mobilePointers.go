package main

import "fmt"

func main() {
	numberOne := 1
	pointerNumberOne := &numberOne
	deName := *pointerNumberOne
	fmt.Println(numberOne)
	fmt.Println(pointerNumberOne)
	fmt.Println(deName)
}
