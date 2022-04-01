package app

func join(segments [][]int64) [][]int64 {
	//Тему с Set.empty на входе я не раскрыл. Не совсем понял требование.
	lenSt := len(segments) - 1
	loop := true
	for loop {
		for i := 0; i < lenSt; i++ {

			// Пузырьковая сортировка по возрастанию начального значения отрезка
			if segments[i][0] > segments[i+1][0] || segments[i][0] == segments[i+1][0] && segments[i][1] > segments[i+1][1] {
				buffer := segments[i]
				segments[i] = segments[i+1]
				segments[i+1] = buffer
				buffer = nil
				break
			}

			// Объединение или поглощение правого отрезка слайса
			if segments[i][1] >= segments[i+1][0] {
				if segments[i][1] >= segments[i+1][1] {
					segments[i][1] = segments[i][1]
				} else {
					segments[i][1] = segments[i+1][1]
				}
				segments = append(segments[:1+i], segments[2+i:]...)
				lenSt--
				if i == lenSt {
					loop = false
				}
				break
			}

			if i == lenSt-1 {
				loop = false
			}

		}
	}

	return segments
}