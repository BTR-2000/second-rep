package main

import (
	"fmt"
	"log"
	"sec/calculate"
	"strings"
)

func main() {
	fmt.Println("Всем салам!")
	readData()
}

func readData() {
	var a, b int
	var op string
	fmt.Print("Введите число A: ")
	if _, err := fmt.Scan(&a); err != nil {
		log.Fatalf("Ошибка! A должно быть целым числом: %v", err)
	}
	fmt.Print("Введите операцию (Варианты: Add, Subtract, Divide, Multiply): ")
	fmt.Scan(&op)
	fmt.Print("Введите число B: ")
	if _, err := fmt.Scan(&b); err != nil {
		log.Fatalf("Ошибка! B должно быть целым числом: %v", err)
	}

	op = strings.TrimSpace(strings.ToLower(op))

	fmt.Print("\nРезультат: ")
	switch op {
	case "add":
		fmt.Printf("%d + %d = %d\n", a, b, calculate.Add(a, b))
	case "subtract":
		fmt.Printf("%d - %d = %d\n", a, b, calculate.Subtract(a, b))
	case "multiply":
		fmt.Printf("%d * %d = %d\n", a, b, calculate.Multiply(a, b))
	case "divide":
		n, err := calculate.Divide(a, b)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			return
		}
		fmt.Printf("%d / %d = %.2f\n", a, b, n)
	default:
		fmt.Println("Не правильно указана операция!")
	}
}
