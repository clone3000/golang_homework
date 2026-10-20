package main

import "fmt"

func printTable(a int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}

func main() {
	var a int
	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	printTable(a)
}
