package main

import "fmt"
import "math"

func main() {
	var a int
	var num int
	maxi := math.MinInt
	mini := math.MaxInt
	fmt.Println("Введите количество чисел: ")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Printf("Введите число %d\t: ", i)
		fmt.Scan(&num)
		if num > maxi {
			maxi = num
		}
		if num < mini {
			mini = num
		}
	}
	fmt.Printf("Максимум : %d\t", maxi)
	fmt.Printf("Минимум : %d\t", mini)
}
