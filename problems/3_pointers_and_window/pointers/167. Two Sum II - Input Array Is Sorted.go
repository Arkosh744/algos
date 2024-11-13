package pointers

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum > target { // если сумма больше таргета, то сдвигаем правый указатель
			r--
			continue
		}
		// иначе сдвигаем левый
		l++
	}

	return []int{-1, -1}
}
