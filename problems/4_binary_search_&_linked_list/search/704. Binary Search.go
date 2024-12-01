package search

/*
https://leetcode.com/problems/binary-search/description/
Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

// time: O(log n)
// mem:  O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)

	for r-l > 1 {
		m := (l + r) / 2
		if good(nums[m], target) {
			l = m
			continue
		}

		r = m
	}

	if nums[l] == target {
		return l
	}

	return -1
}

// < если ответ в правом указателе
// <= если ответ в левом
// меняем только функцию good и то, где лежит ответ

func good(val, target int) bool {
	return val <= target
}
