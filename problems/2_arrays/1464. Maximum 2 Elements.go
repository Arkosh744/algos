package arrays

// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

func maxProduct(nums []int) int {
	// [3,4,5,2]
	var num1, num2 int
	for i := range nums {
		if nums[i] > num1 {
			num2 = num1
			num1 = nums[i]
			continue
		}

		if nums[i] > num2 {
			num2 = nums[i]
		}
	}

	return (num1 - 1) * (num2 - 1)
}

/*
Оценка по времени: О(n)
Проходим по массиву 1 раз, где n- количество элементов в массиве

Оценка по памяти: О(1)
Создаем всегда 2 переменных, независимо от количества входных данных

Описание решения.
Заводим 2 переменных и проходим по массиву двигая указатель.
Если 1 число меньше текущего, то переносим его во 2 и записываем новый максимум.
Так же не забываем проверять 2 число.
Возвращаем результат по формуле
*/
