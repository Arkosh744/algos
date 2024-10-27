package retry

// hashMap
/*
Given n points on a 2D plane, find if there is such a line parallel to the y-axis that reflects the given points
symmetrically.
In other words, answer whether or not if there exists a line that after reflecting all points over the given line, the
original points' set is the same as the reflected ones.
Note that there can be repeated points.
*/
// https://leetcode.com/problems/line-reflection/description/

func isReflected(points [][]int) bool {
	// находим минимальный и максимальный X
	maxX := points[0][0]
	minX := points[0][0]
	for _, point := range points {
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[0] < minX {
			minX = point[0]
		}
	}

	// сделали хеш-таблицу, чтобы проверять наличие точки за O(1)
	pointSet := make(map[[2]int]bool)
	for _, point := range points {
		pointSet[[2]int{point[0], point[1]}] = true
	}

	for _, point := range points {
		// symX = maxX + minX - x
		// symX - координата точки по оси X симметричной x
		symX := maxX + minX - point[0]
		if !pointSet[[2]int{symX, point[1]}] {
			return false
		}
	}

	return true
}
