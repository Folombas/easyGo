package main

import "fmt"

func main() {
	// Цены на чай и кофе в зависимости от размера
	teaPrices := map[string]float64{
		"маленький": 50,
		"средний":   70,
		"большой":   90,
	}
	coffeePrices := map[string]float64{
		"маленький": 80,
		"средний":   100,
		"большой":   120,
	}

	// Доступные размеры
	sizes := []string{"маленький", "средний", "большой"}

	fmt.Println("Добро пожаловать в кофейню!")
	fmt.Println("Выберите размер стакана для чая:")
	for i, size := range sizes {
		fmt.Printf("%d. %s\n", i+1, size)
	}
	var teaChoice int
	fmt.Scan(&teaChoice)
	if teaChoice < 1 || teaChoice > 3 {
		fmt.Println("Некорректный выбор, установлен 'маленький'")
		teaChoice = 1
	}
	teaSize := sizes[teaChoice-1]
	teaPrice := teaPrices[teaSize]

	fmt.Println("Выберите размер стакана для кофе:")
	for i, size := range sizes {
		fmt.Printf("%d. %s\n", i+1, size)
	}
	var coffeeChoice int
	fmt.Scan(&coffeeChoice)
	if coffeeChoice < 1 || coffeeChoice > 3 {
		fmt.Println("Некорректный выбор, установлен 'маленький'")
		coffeeChoice = 1
	}
	coffeeSize := sizes[coffeeChoice-1]
	coffeePrice := coffeePrices[coffeeSize]

	// Вывод информации
	fmt.Println("\n--- Ваш заказ ---")
	fmt.Printf("%s чай стоит: %.2f руб.\n", teaSize, teaPrice)
	fmt.Printf("%s кофе стоит: %.2f руб.\n", coffeeSize, coffeePrice)
	fmt.Printf("Итого: %.2f руб.\n", teaPrice+coffeePrice)
}
