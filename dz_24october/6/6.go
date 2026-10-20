package main

import "fmt"

func main() {

	var n int
	var num int
	plus := 0
	minus := 0
	zero := 0

	fmt.Println("Введите количество чисел: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Введите число %d\t: ", i)
		fmt.Scan(&num)

		if num > 0 {
			plus += 1
		} else if num < 0 {
			minus += 1
		} else {
			zero += 1
		}
	}
	fmt.Printf("Положительных: %d Отрицательных: %d Нулей: %d", plus, minus, zero)
}
