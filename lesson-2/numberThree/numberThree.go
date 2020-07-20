package main

import (
	"fmt"
)

func numberThree(number int) {
	if number%3 == 0 {
		fmt.Printf("Число %v делится на 3 без остатка \n", number)
	} else {
		fmt.Printf("Число %v не делится на 3 без остатка \n", number)
	}
}

func main() {
	var number int
	fmt.Println("Введите число для проверки целочисленного деления на три")
	fmt.Scanln(&number)
	numberThree(number)
}
