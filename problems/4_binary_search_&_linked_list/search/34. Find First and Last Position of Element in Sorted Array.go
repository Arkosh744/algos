package search

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := searchFirst(nums, target)
	last := searchLast(nums, target)

	return []int{first, last}
}

func searchFirst(nums []int, target int) int {
	l, r := -1, len(nums)-1

	for r-l > 1 {
		m := (l + r) / 2
		if goodFirst(nums[m], target) {
			l = m
			continue
		}

		r = m
	}

	if nums[r] == target {
		return r
	}

	return -1
}

func searchLast(nums []int, target int) int {
	l, r := 0, len(nums)

	for r-l > 1 {
		m := (l + r) / 2
		if goodLast(nums[m], target) {
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
// меняем только функцию good и то, где лежит ответ (в строке nums[l] или nums[r])

func goodLast(val, target int) bool {
	return val <= target
}

func goodFirst(val, target int) bool {
	return val < target
}
