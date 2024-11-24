package window

/* https://leetcode.com/problems/max-consecutive-ones/description/
Example 1:

Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
Example 2:

Input: nums = [1,0,1,1,0,1]
Output: 2
*/

func findMaxConsecutiveOnes(nums []int) int {
	var l, r int
	var maxCount int

	for l < len(nums) {
		for r+1 < len(nums) && nums[r] == nums[r+1] {
			r += 1
		}

		if nums[r] > 0 {
			maxCount = max(maxCount, r-l+1)
		}

		l = r + 1
		r = r + 1
	}

	return maxCount
}
