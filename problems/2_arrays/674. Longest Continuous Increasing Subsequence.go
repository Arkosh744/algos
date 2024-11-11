package arrays

import "math"

// https://leetcode.com/problems/longest-continuous-increasing-subsequence/

func findLengthOfLCIS(nums []int) int {
	var count, maxCount int
	lastNumber := math.MinInt
	for _, num := range nums {
		if lastNumber < num {
			count++
		} else {
			count = 1
		}
		maxCount = max(maxCount, count)
		lastNumber = num
	}

	return maxCount
}
