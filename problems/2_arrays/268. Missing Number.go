package arrays

// https://leetcode.com/problems/missing-number/description/

func missingNumber(nums []int) int {

	var res int
	for i := 0; i <= len(nums); i++ {
		res += i
	}

	for _, n := range nums {
		res -= n
	}

	return res
}
