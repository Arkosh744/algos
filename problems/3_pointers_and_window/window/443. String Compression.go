package window

import "strconv"

func compress(chars []byte) int {
	l, r := 0, 0

	for r < len(chars) {
		start := r // Начало окна в текущей группе сжатия
		// Пробегаем по всем одинаковым символам
		for r < len(chars) && chars[r] == chars[start] {
			r++
		}

		// Записываем символ
		chars[l] = chars[start]
		l++

		// Если символы повторялись, записываем количество
		count := r - start // Вычисляем количество повторов
		if count > 1 {
			countStr := strconv.Itoa(count)
			for i := 0; i < len(countStr); i++ {
				chars[l] = countStr[i]
				l++
			}
		}
	}

	return l
}
