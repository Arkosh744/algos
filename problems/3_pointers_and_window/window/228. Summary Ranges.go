package window

import "fmt"

/*
https://leetcode.com/problems/summary-ranges/description/

Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

*/

func summaryRanges(nums []int) []string {
	var l, r int
	var result []string

	for l < len(nums) {
		for r+1 < len(nums) && nums[r]+1 == nums[r+1] {
			r += 1
		}

		// добавляем в ответ
		if r != l {
			result = append(result, fmt.Sprintf("%d->%d", nums[l], nums[r]))
		} else {
			result = append(result, fmt.Sprintf("%d", nums[l]))
		}

		// интервалы не пересекаются, поэтому сдвигаем
		// на r + 1 - именно отсюда будет начинаться
		// следующий интервал
		l = r + 1
		r = r + 1
	}

	return result
}
