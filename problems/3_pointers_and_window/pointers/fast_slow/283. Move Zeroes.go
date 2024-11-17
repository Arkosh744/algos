package fast_slow

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// https://leetcode.com/problems/move-zeroes/description/

func moveZeroes(nums []int) {
	slow := 0

	for _, num := range nums { // быстрым указателем ищем не нулевой элемент
		if num == 0 {
			continue
		}
		nums[slow] = num // как нашли - подставляем его в slow
		slow++           // запоминаем позицию после которой сделаем все нулями
	}

	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
