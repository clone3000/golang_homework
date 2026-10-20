package main

import "fmt"
import "math"

func main() {
	var a int
	var num int
	var total int
	var average float64
	var count int
	var nocount int

	maxi := math.MinInt
	mini := math.MaxInt
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		fmt.Printf("Введите число %d\t: ", i)
		fmt.Scan(&num)

		total = total + num

		if num > maxi {
			maxi = num
		}
		if num < mini {
			mini = num
		}

		if num%2 == 0 {
			count += 1
		} else {
			nocount += 1
		}

	}
	average = float64(total) / float64(a)

	fmt.Printf("Сумма чисел: %d\t", total)
	fmt.Printf("Среднее значение: %v\t", average)
	fmt.Printf("Максимум: %d\t", maxi)
	fmt.Printf("Минимум: %d\t", mini)
	fmt.Printf("Четных: %d\t", count)
	fmt.Printf("Нечетных: %d\t", nocount)
}
