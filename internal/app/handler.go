package app

import (
	"bufio"
	"fmt"
	"os"
)

func Handler(funcNum int) {
	switch funcNum {
	case 1:
		stdin := bufio.NewScanner(os.Stdin)
		fmt.Printf("Введите произвольный набор символов АНГЛИЙСКОГО алфавита на вход rle() функции: ")
		stdin.Scan()
		if err := stdin.Err(); err != nil || !isEngChr(stdin.Text()) {
			Handler(funcNum)
			return
		}
		fmt.Printf("Результирующая закодированная строка: ")
		fmt.Println(rle(stdin.Text()))
		return

	case 2:
		stdin := bufio.NewReader(os.Stdin)
		fmt.Printf("Введите набор отрезков чисел на вход join() функции, например `(1,2),(5,100),(2,3)` или `(-5,10),(15,20),(10,25),(100,200)`: ")
		var str string
		_, _ = fmt.Fscan(stdin, &str)
		segments := strToSegments(str)
		resSegments := join(segments)
		if len(resSegments) > 0 {
			fmt.Print("Результирующий набор отрезков: ")
			fmt.Println(resSegments)
		}
		return

	}
}
