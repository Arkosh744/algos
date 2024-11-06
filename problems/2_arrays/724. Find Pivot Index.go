package arrays

// https://leetcode.com/problems/find-pivot-index/description/

func pivotIndex(nums []int) int {
	// [1,7,3,6,5,6]

	var sum int // 28
	for _, n := range nums {
		sum += n
	}

	leftSum := 0
	rightSum := sum // 28
	for i, n := range nums {
		rightSum -= n                // 27 20 17 11
		leftSum = sum - n - rightSum // 0 1 8 11
		if rightSum == leftSum {
			return i
		}
	}

	return -1
}

/*
Оценка по времени: О(n)
Проходим по массиву 2 раза, где n- количество элементов в массиве

Оценка по памяти: О(1)
Создаем всегда 3 переменных, независимо от количества входных данных

Описание решения.
Вначале находим сумму всех элементов массива.
Далее двигая указатель находим сумму левой и правой части массивов, не забывая вычитать текущий элемент из обоих.
*/
