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

	sizes := []string{"маленький", "средний", "большой"}

	fmt.Println("Добро пожаловать в кофейню!")
	fmt.Println("Что вы хотите заказать?")
	fmt.Println("1. Чай")
	fmt.Println("2. Кофе")
	fmt.Println("3. И чай, и кофе")
	var choice int
	fmt.Scan(&choice)

	var teaSize string
	var coffeeSize string
	var teaPrice, coffeePrice float64

	switch choice {
	case 1:
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
		teaSize = sizes[teaChoice-1]
		teaPrice = teaPrices[teaSize]
		coffeePrice = 0
	case 2:
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
		coffeeSize = sizes[coffeeChoice-1]
		coffeePrice = coffeePrices[coffeeSize]
		teaPrice = 0
	case 3:
		fmt.Println("Выберите размер стакана для чая:")
		for i, size := range sizes {
			fmt.Printf("%d. %s\n", i+1, size)
		}
		var teaChoice int
		fmt.Scan(&teaChoice)
		if teaChoice < 1 || teaChoice > 3 {
			fmt.Println("Некорректный выбор для чая, установлен 'маленький'")
			teaChoice = 1
		}
		teaSize = sizes[teaChoice-1]
		teaPrice = teaPrices[teaSize]

		fmt.Println("Выберите размер стакана для кофе:")
		for i, size := range sizes {
			fmt.Printf("%d. %s\n", i+1, size)
		}
		var coffeeChoice int
		fmt.Scan(&coffeeChoice)
		if coffeeChoice < 1 || coffeeChoice > 3 {
			fmt.Println("Некорректный выбор для кофе, установлен 'маленький'")
			coffeeChoice = 1
		}
		coffeeSize = sizes[coffeeChoice-1]
		coffeePrice = coffeePrices[coffeeSize]
	default:
		fmt.Println("Некорректный выбор, попробуйте снова.")
		return
	}

	// Вывод информации
	fmt.Println("\n--- Ваш заказ ---")
	if teaPrice > 0 {
		fmt.Printf("%s чай стоит: %.2f руб.\n", teaSize, teaPrice)
	}
	if coffeePrice > 0 {
		fmt.Printf("%s кофе стоит: %.2f руб.\n", coffeeSize, coffeePrice)
	}
	fmt.Printf("Итого: %.2f руб.\n", teaPrice+coffeePrice)
}