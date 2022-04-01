package app

import (
	"strconv"
	"strings"
)

func isEngChr(input string) bool {
	for i := 0; i < len(input); i++ {
		if (input[i] >= 'a' && input[i] <= 'z') || (input[i] >= 'A' && input[i] <= 'Z') {
			continue
		} else {
			return false
		}
	}

	return true
}

func strToSegments(str string) [][]int64 {
	inputStr := str
	inputStr = strings.TrimLeft(inputStr, "(")
	inputStr = strings.TrimRight(inputStr, ")")
	dataSet := strings.Split(inputStr, "),(")
	lenSet := len(dataSet)
	lenCup := cap(dataSet)

	//Заглушка от паники. Детальную валидацию входных данных делать не стал.
	if lenSet == 1 {
		return make([][]int64, 1, 1)
	}

	segments := make([][]int64, lenSet, lenCup)
	for i := 0; i < lenSet; i++ {
		elems := strings.Split(dataSet[i], ",")
		x0, _ := strconv.ParseInt(elems[0], 10, 64)
		x1, _ := strconv.ParseInt(elems[1], 10, 64)
		segments[i] = []int64{x0, x1}
	}

	return segments
}
