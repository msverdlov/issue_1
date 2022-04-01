package app

import (
	"strconv"
)

func rle(s string) string {
	if s == "" {
		return ""
	}

	sLen := len(s) - 1
	lenCurVector := 1
	result := ""

	for i := 0; i < sLen; i++ {
		if s[i] == s[i+1] {
			lenCurVector++
		} else {
			result += string(s[i])
			if lenCurVector > 1 {
				result += strconv.Itoa(lenCurVector)
			}
			lenCurVector = 1
		}
	}

	result += string(s[sLen])
	if lenCurVector > 1 {
		result += strconv.Itoa(lenCurVector)
	}

	return result
}
