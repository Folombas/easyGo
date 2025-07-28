package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID      int     `json:"id" max:"100"`
	Name    string  `json:"name" required:"true"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance,omitempty"`
}

func main() {
	// 1. Проверка типа и значения
	num := 42
	fmt.Println("Тип num:", reflect.TypeOf(num))       // int
	fmt.Println("Значение num:", reflect.ValueOf(num)) // 42

	// 2. Исследование структуры
	user := User{
		ID:    1,
		Name:  "Gosha",
		Email: "gosha@example.com",
	}

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	// 2.1. Количество полей
	fmt.Println("\nКоличество полей:", t.NumField())

	// 2.2. Итерация по полям
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("\nПоле %d: %s\n", i+1, field.Name)
		fmt.Printf("  Тип: %s\n", field.Type)
		fmt.Printf("  Значение: %v\n", value.Interface())
		fmt.Printf("  JSON-тег: %s\n", field.Tag.Get("json"))

		if req, ok := field.Tag.Lookup("required"); ok {
			fmt.Printf("  Обязательное: %s\n", req)
		}
	}

	// 3. Изменение значения через указатель
	ptr := &user
	vPtr := reflect.ValueOf(ptr).Elem()
	balanceField := vPtr.FieldByName("Balance")

	if balanceField.CanSet() {
		balanceField.SetFloat(99.99)
		fmt.Println("\nНовый баланс:", user.Balance) // 99.99
	}

	// 4. Динамический вызов функции
	funcValue := reflect.ValueOf(greet)
	args := []reflect.Value{reflect.ValueOf("Gosha")}
	funcValue.Call(args) // Привет, Gosha!
}

func greet(name string) {
	fmt.Printf("\nПривет, %s!\n", name)
}
