package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите N: ")
	fmt.Scan(&n)
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Printf("Сумма всех чисел: %d Среднее значение %v\n", total, float64(total)/float64(n))
}
