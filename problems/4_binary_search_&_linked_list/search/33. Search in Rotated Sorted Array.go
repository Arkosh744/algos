package search

// time: O(log n)
// mem:  O(1)
func findOffset(nums []int) int {
	// good   bad               //  good        bad
	// [   |  1 2 3 4 5]        // [4 5 6 7  |  0 1 2]
	//   l    r                 //        l     r

	good := func(i int) bool {
		return nums[i] > nums[len(nums)-1]
	}

	l, r := -1, len(nums)-1
	for r-l > 1 {
		m := (l + r) / 2
		if good(m) {
			l = m
			continue
		}

		r = m
	}

	return r
}

func search(nums []int, target int) int {
	good := func(i int) bool {
		return nums[i] <= target
	}

	// обычный бинарный поиск, но смещаем на offset дополнительно
	offset := findOffset(nums)
	l, r := offset, len(nums)+offset
	for r-l > 1 {
		// Note: ошибка №1 это делать "m = (l + r + offset) // 2"
		m := (l + r) / 2
		if good(m % len(nums)) {
			l = m
			continue
		}

		r = m
	}

	// Note: ошибка №2 это забыть сделать "(l + offset) % len(nums)"
	realLeft := l % len(nums)
	if nums[realLeft] == target {
		return realLeft
	}

	return -1
}
