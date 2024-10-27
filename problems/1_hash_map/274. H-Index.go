package hashMap

// https://leetcode.com/problems/h-index/

func hIndex(citations []int) int {
	// 10, 1, 8, 0, 3
	arr := make([]int, len(citations)+1)
	for _, val := range citations {
		if val >= len(citations) {
			val = len(citations)
		}

		arr[val]++
	}

	// [1 1 0 1 0 2]
	var papersCount int
	for hResult := len(citations); hResult >= 0; hResult-- {
		papersCount += arr[hResult]

		if papersCount >= hResult {
			return hResult
		}
	}

	return 0
}

/*
Оценка по времени: О(n)
Сначала проходим массив для подсчёта потом проходим с поиском, в худшем случае 2n итераций -> O(n)

Оценка по памяти: О(n)
Выделяем памяти для массива как n+1

Описание решения
Сначала посчитаем сколько раз, но не больше n цитировалась каждая статься и запишем на место минимум из количества цитирований
и максимально возможный H-индекс равный `len(citations)`
Потом с конца массива идём и плюсуем значения пока индекс не будет меньше или равен количеству статей.
Проходим до первого, если ничего не нашли возвращаем 0
*/
