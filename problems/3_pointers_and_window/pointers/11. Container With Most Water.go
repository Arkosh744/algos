package pointers

import "fmt"

func maxArea(height []int) int {
	var area int
	p1 := 0
	p2 := len(height) - 1

	for p1 < p2 {
		fmt.Println(height[p1], height[p2])
		fmt.Println(p1, p2)
		curArea := min(height[p1], height[p2]) * (p2 - p1)
		fmt.Println(curArea)

		area = max(area, curArea)

		if height[p2] > height[p1] {
			p1++
		} else {
			p2--
		}
	}

	return area
}
