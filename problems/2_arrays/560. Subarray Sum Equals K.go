package arrays

// https://leetcode.com/problems/subarray-sum-equals-k/description/

func subarraySum(nums []int, target int) int {
	prefixSums := map[int]int{0: 1} // ключ - префиксная сумма, значение - сколько раз встретили
	currentPrefixSum := 0           // текущая префиксная сумма
	resCount := 0                   // ответ

	for _, el := range nums {
		currentPrefixSum += el

		// проверяем встречали ли мы уже префиксный массив с суммой current_prefix_sum - target
		if _, ok := prefixSums[currentPrefixSum-target]; ok {
			// если встречали - то к ответу прибавляем число сколько раз уже встретили
			resCount += prefixSums[currentPrefixSum-target]
		}

		// добавляем текущую префиксную сумму в массив
		if _, ok := prefixSums[currentPrefixSum]; !ok {
			prefixSums[currentPrefixSum] = 0
		}
		prefixSums[currentPrefixSum] += 1
	}

	return resCount
}

/*
Мы заранее создали хеш-таблицу, в которой будем хранить в качестве ключа значение сумм (префиксная хеш-таблица, как префиксный массив).
Сразу заполнили ключом 0 и значением 1 (сумма пустого массива).

Каждую итерацию мы currentPrefixSum увеличиваем на nums[i] и
добавляем в конце хода получившуюся currentPrefixSum в хэш-таблицу со значением 1, если такого ключа ранее не было, либо с value+1 для существующего ключа.
Перед добавлением currentPrefixSum в хэш-таблицу проверяем наличие ключа = currentPrefixSum - target,
если находим, то берём значение и добавляем к итоговому ответу (т.е. количество отрезков, сумма элементов в котором = значению currentPrefixSum - target)
возможно стоит пересмотреть https://kinescope.io/jdxLNhW7RMyQRPRGPwpqCc
*/
