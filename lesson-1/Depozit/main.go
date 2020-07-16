package main

import (
	"fmt"
)

func main() {
	var Summa float64
	fmt.Println("Введите размер вклада")
	fmt.Scanln(&Summa)
	var Procent float64
	fmt.Println("Введите размер годовых процентов по вкладу")
	fmt.Scanln(&Procent)
	Summa = Summa * (100 + Procent) / 100
	Summa = Summa * (100 + Procent) / 100
	Summa = Summa * (100 + Procent) / 100
	Summa = Summa * (100 + Procent) / 100
	Summa = Summa * (100 + Procent) / 100
	fmt.Println("Через пять лет сумма вклада составит ", Summa)
}
