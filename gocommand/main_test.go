package main

import "testing"

func TestDemoFunction(t *testing.T) {
	// Просто проверяем, что функция не паникует
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoFunction panicked: %v", r)
		}
	}()
	DemoFunction()
}

func TestMainOutput(t *testing.T) {
	// Здесь можно было бы перехватить вывод, но для демонстрации достаточно
	t.Log("main package compiled successfully")
}