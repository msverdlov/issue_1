package main

import (
	"bufio"
	"fmt"
	"issue_1/internal/app"
	"os"
)

const countFunc = 2

func main() {
	hello()
	handler()
	bye()
}

func hello() {
	fmt.Println("Привет!")
	fmt.Printf("Данное приложение реализует %d функции:\n", countFunc)
	fmt.Println("1. Run-length encoding")
	fmt.Println("2. Объединение отрезков")
}

func handler() {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Printf("Введите номер функции для запуска от 1 до %d : ", countFunc)

	var funNum int
	_, err := fmt.Fscan(stdin, &funNum)
	if err != nil || funNum <= 0 || funNum > countFunc {
		handler()
		return
	}

	app.Handler(funNum)
}

func bye() {
	fmt.Println("Спасибо, программа завершена! Пока!")
}
