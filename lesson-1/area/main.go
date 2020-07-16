package main

import (
	"fmt"
	"math"
)

func main() {
	var Katet1 float64
	fmt.Println("Выедите длину первого катета")
	fmt.Scanln(&Katet1)
	var Katet2 float64
	fmt.Println("Введите длину второго катета")
	fmt.Scanln(&Katet2)
	var Area float64
	Area = Katet1 * Katet2 / 2
	fmt.Println("Площадь прямоугольного треугольника = ", Area)
	var Gipotenuza float64
	Gipotenuza = math.Sqrt((Katet1 * Katet1) + (Katet2 * Katet2))
	fmt.Println("Гипотенуза прямоугольного треугольника = ", Gipotenuza)
	var Perimeter float64
	Perimeter = Katet1 + Katet2 + Gipotenuza
	fmt.Println("Периметр прямоугольного треугольника = ", Perimeter)
}
