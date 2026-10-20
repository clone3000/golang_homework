package main

import "fmt"

func main() {
	var a float64
	var operation string
	var b float64
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Операция (+, -, *, /): ")
	fmt.Scan(&operation)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя")
		} else {

			fmt.Println(a / b)
		}

	default:
		fmt.Println("Неверная операция")
	}

}
