package hashMap

// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/

func findThePrefixCommonArray(A []int, B []int) []int {
	if len(A) != len(B) {
		return nil
	}

	arr := make([]int, len(A)+1)
	var res []int

	var count int
	for i := 0; i < len(A); i++ {
		arr[A[i]]++
		arr[B[i]]++

		if arr[A[i]] == 2 || arr[B[i]] == 2 {
			count++
		}

		res = append(res, count)
	}

	return res
}

/*
Оценка по времени: О(n)
Проходим по массиву 1 раз

Оценка по памяти: О(n)
Выделяем памяти для массива как n

Описание решения.
Создаём массив из n+1 и счётчик, проходим по двум массивам параллельно.
При проходе +1 к индексу который равен значению элемента.
Проверяем если в массиве элементы стали равны 2, то +1 к счётчику и добавляем счётчик в массив ответа
*/
