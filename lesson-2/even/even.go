package main

import (
	"fmt"
)

func even(number int) {
	if number%2 == 0 {
		fmt.Printf("Число %v чётное \n", number)
	} else {
		fmt.Printf("Число %v нечётное \n", number)
	}
}

func main() {
	var number int
	fmt.Println("Введите число, чётность которого вы хоите определить")
	fmt.Scanln(&number)
	even(number)
}