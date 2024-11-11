package arrays

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/

// time O(n logN) - зависит от сортировки и ее реализации
// mem O(1) - аллоцируем всегда 3 переменных
func minimumDifference(nums []int, k int) int {
	result := math.MaxInt

	sort.Ints(nums) // сортируем

	// проходимся по сортированному массиву находя разницу между левым и правым указателем, длина определяется k
	pointer1 := 0
	pointer2 := k - 1

	for pointer2 < len(nums) {
		result = min(result, nums[pointer2]-nums[pointer1])
		pointer1++
		pointer2++
	}

	return result
}
