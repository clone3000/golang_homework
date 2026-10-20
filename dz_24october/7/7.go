package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите N: ")
	fmt.Scan(&n)
	total := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			total += i
		}

	}
	fmt.Println(total)
}
