package main

import "fmt"

func main() {
	var name string = "Gosha"
	ptrName := &name
	deName := *ptrName
	fmt.Println(name)
	fmt.Println(ptrName)
	fmt.Println(deName)
}
