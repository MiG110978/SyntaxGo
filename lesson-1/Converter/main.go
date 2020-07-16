package main

import (
	"fmt"
)

func main() {
	const UsdRub float64 = 70
	var Rub float64
	fmt.Println("Введите сумму в рублях для пересчета в доллары")
	fmt.Scanln(&Rub)
	var Usd float64
	Usd = Rub / UsdRub
	fmt.Println(Rub, " рублей = ", Usd, " долларов.")
}
