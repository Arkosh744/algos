package pointers

import (
	"math"
)

// https://leetcode.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
	// Input: nums = [-4,-1,0,3,10]
	// Output: [0,1,9,16,100]

	res := make([]int, len(nums)) // будем хранить массив и перезаписывать в убывающем порядке

	p1 := 0
	p2 := len(nums) - 1

	for i := p2; i >= 0; i-- { // нам нужно точно заполнить массив из всех элементов поэтому условие >= 0
		if abs(nums, p1) > abs(nums, p2) { // выбираем с какой стороны подвинем указатель
			res[i] = nums[p1] * nums[p1]
			p1++
			continue
		}

		res[i] = nums[p2] * nums[p2]
		p2--
	}

	return res
}

func abs(nums []int, p1 int) int {
	return int(math.Abs(float64(nums[p1])))
}
