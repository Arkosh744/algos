package arrays

// https://leetcode.com/problems/rotate-array/

// time O(n) полностью проходим массив
// mem O(1) ничего нового не создаем
func rotate(nums []int, k int) {
	k = k % len(nums)

	rotateSubArr(nums, 0, len(nums)) // переворачиваем полностью
	rotateSubArr(nums, 0, k)         // переворачиваем от 0 до k, не включая позицию k
	rotateSubArr(nums, k, len(nums)) // переворачиваем оставшуюся часть
}

func rotateSubArr(nums []int, i, j int) {
	// nums=[1, 2, 3, 4], i=1, j=3 -> [1, 3, 2, 4]
	j--
	// i=1, j=2
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
