package hashMap

// https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		hashMap[v]++
	}

	// индекс массива - сколько раз встретилось число
	// значение - список чисел, которые встретились столько раз
	frequencyList := make([][]int, len(nums)+1)
	for num, frequency := range hashMap {
		frequencyList[frequency] = append(frequencyList[frequency], num)
	}

	// допустим у нас получился такой frequencyList:
	// 0: []
	// 1: [2, 5]
	// 2: []
	// 3: [4]
	// 4: []
	// 5: []
	// при k = 1 нам нужно вернуть 4
	// для этого проходимся с конца и ищем первые k элементов
	var result []int
	for i := len(frequencyList) - 1; i >= 0; i-- {
		for _, num := range frequencyList[i] {
			if k <= 0 {
				return result
			}

			result = append(result, num)
			k--
		}
	}

	return result
}

/*
Оценка по времени: О(n)
Проходим по массиву 1 раз потом ещё раз по результирующему всего 2n проходов

Оценка по памяти: О(n)
Выделяем памяти для мамы и для массива в зависимости от n

Описание решения.
1. создаем мапу счетчик
2. создаём срез срезов длиной как количество элементов +1 и кладём туда элементы, где индекс это количество вхождений
3. Проходим с конца этого среза и забираем элементы пока не получим k значений
*/
