package main

import "fmt"

func main() {
	// Явное объявление переменных
	var teaName string = "Green Tea"
	var hotCoffee string = "Hot coffee"

	// Неявное объявление переменных
	var size = "Little"
	var sizeCoffee = "Middle"

	var price = 50.00
	var priceCoffee = 100.00

	fmt.Println("Little Green Tea price is ₽50.00")
	fmt.Println(size, teaName, "price is ₽", price)
}
