package pointersnice

import "fmt"

func main() {
	x := 37   // Обычная переменная
	ptr := &x // Указатель на x (адрес x)

	fmt.Println(x)    // 37 - значение переменной x
	fmt.Println(ptr)  // 0xc000... - адрес переменной x (например, 0xc000014078)
	fmt.Println(*ptr) // 37 - разыменовываем ptr -> получаем значение, на которое он указывает
}
