package main

import (
	"fmt"
	"strconv"
)

var firstNum float32
var secondNum float32
var action string

func calculate(firstNum float32, action string, secondNum float32) float32 {
	switch action {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	default:
		fmt.Println("Не удалось распознать действие")
		return 0
	}
}

func inputData(msg string) (data string) {
	fmt.Print(msg)
	fmt.Scanln(&data)
	//if data == "exit" {
	//	panic(nil)
	//}
	switch data {
	case "exit":
		panic(nil)
	case "help":
		fmt.Printf("В текущей версии калькулятора реализванны функции:\n сложения(вводится +), \n вычитания (вводится -), \n умножения (вводится *),\n деления (вводится /)")
		panic(nil)
	}
	return
}

func enterNumber(msg string) float32 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			return float32(num)
			break
		}
		fmt.Println("Не удалось распознать число")
	}
	return 0
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// Ввод первого числа
	firstNum := enterNumber("Введите первое число: ")

	action := inputData("Укажите действие: ")

	// Ввод второго числа
	secondNum := enterNumber("Введите второе число: ")

	fmt.Printf("Результат: %v \n", calculate(firstNum, action, secondNum))

}
