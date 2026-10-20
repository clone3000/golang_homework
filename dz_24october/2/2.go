package main

import "fmt"

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Положительное")
	} else if a < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}

	if a%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}

	if a%3 == 0 && a%5 == 0 {
		fmt.Println("Делимое")
	} else {
		fmt.Println("Неделимое")
	}
}
