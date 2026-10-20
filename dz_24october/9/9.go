package main

import "fmt"

func main() {
	fmt.Println("Выберете операцию:\n 1 - Посмотреть баланс\n 2 - Пополнение \n 3 - Снять наличные \n 4 - Выход")
	balance := 0

	for {
		var operation int
		fmt.Scan(&operation)
		switch operation {
		case 1:
			fmt.Printf("Баланс: %d рублей\n", balance)
		case 2:
			var popolnenie int
			fmt.Print("Введите сумму: ")
			fmt.Scan(&popolnenie)
			balance = popolnenie + balance
			fmt.Printf("Зачислено: %d рублей\n", popolnenie)
		case 3:
			var snyatie int
			fmt.Print("Введите сумму: ")
			fmt.Scan(&snyatie)
			if snyatie > balance {
				println("Недостаточно средств")
			} else {
				balance = balance - snyatie
				fmt.Printf("Снятие: %d рублей\n", snyatie)
			}
		case 4:
			fmt.Println("Выход из системы")
			return

		default:
			fmt.Println("Неизвестная операция")
		}
	}
}
