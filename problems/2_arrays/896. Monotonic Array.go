package arrays

// https://leetcode.com/problems/monotonic-array/

func isMonotonic(nums []int) bool {
	var (
		isInc = true
		isDec = true
	)

	for i := 1; i < len(nums); i++ {
		isInc = isInc && nums[i-1] <= nums[i]
		isDec = isDec && nums[i-1] >= nums[i]
	}

	return isInc || isDec
}

/*
Оценка по времени: О(n)
Проходим по массиву 1 раз, где n- количество элементов в массиве

Оценка по памяти: О(1)
Создаем всегда 2 переменных, независимо от количества входных данных

Описание решения.
на каждом участке массива проверяем выполнение условия и меняем флаг
в конце возвращаем "или" из обоих флагов
*/
