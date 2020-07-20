package main

import (
	"fmt"
	"math/big"
)

func fibonachi100() {
	fibonachi2 := big.NewInt(0)
	fmt.Println(fibonachi2)
	fibonachi1 := big.NewInt(1)
	fmt.Println(fibonachi1)
	fibonachi := big.NewInt(0)
	for i := 3; i <= 100; i++ {
		fibonachi.Add(fibonachi1,fibonachi2)
		fmt.Println(fibonachi)
		fibonachi2 = fibonachi1
		fibonachi1 = fibonachi
	}
}

func main() {
	fibonachi100()
}
