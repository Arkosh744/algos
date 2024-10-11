package hashMap

// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int { // нужно вернуть индексы слагаемых
	numMap := make(map[int]int)

	for idx, val := range nums {
		if position, ok := numMap[target-val]; ok {
			return []int{idx, position}
		}

		numMap[val] = idx
	}

	return nil
}

/*
Оценка по времени: О(n)
Проходим по массиву 1 раз, где n- количество элементов в массиве

Оценка по памяти: О(n)
Выделяем дополнительную память под мапу в зависимости от размера входного массива

Описание решения.
Чтобы найти пару значений можно искать таргет минус текущее число в мапе,
в которую мы на каждом шаге сохраняем число и его позицию.
Если такое число есть - возвращаем текущую позицию и ключ нашей мапы.
*/
