
package main

import "fmt"

func main() {
	name := "Green Tea Flying Dragon"
	price := 2.55 // default is float64
	ready := true
	shopCount := 8
	var stockCount int64 = 2000

	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of price is: %T\n", price)
	fmt.Printf("Type of ready is: %T\n", ready)
	fmt.Printf("Type of shopCount is: %T\n", shopCount)
	fmt.Printf("Type of stockCount is: %T\n", stockCount)
}
