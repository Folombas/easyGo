package main

import "fmt"

func main() {
	// Явное объявление переменных
	var teaName string = "Green Tea"

	// Неявное объявление переменных
	var size = "Little"

	var price = 50.00

	fmt.Println("Little Green Tea price is ₽50.00")
	fmt.Println(size, teaName, "price is ₽", price)
}
