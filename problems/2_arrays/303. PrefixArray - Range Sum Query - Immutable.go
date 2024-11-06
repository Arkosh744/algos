package arrays

// https://leetcode.com/problems/range-sum-query-immutable/description/

type NumArray struct {
	arr []int
}

// time: O(n)
// mem: O(n)

func Constructor(nums []int) NumArray {
	sumArr := make([]int, len(nums)+1)
	// [1, 2, 3] -> [0, 1, 3, 6] собираем префиксный массив
	for i := 0; i < len(nums); i++ {
		sumArr[i+1] = nums[i] + sumArr[i]
	}

	return NumArray{arr: sumArr}
}

// time: O(1)
// mem: O(1)

func (a *NumArray) SumRange(left int, right int) int {
	return a.arr[right+1] - a.arr[left]
}

/*
Описание решения.
В конструкторе собираем префиксный массив сумм элементов, добавляя вначале нулевой элемент
Далее выдаем нужную сумму с левого по правый элемент вычитая из одного значения массива другое по индексам
*/
