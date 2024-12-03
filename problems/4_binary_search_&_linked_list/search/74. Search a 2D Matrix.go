package search

// https://leetcode.com/problems/search-a-2d-matrix/description/
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

func searchMatrix(matrix [][]int, target int) bool {
	matrixLen := len(matrix[0])

	matrixToRow := func(i int) int {
		return matrix[i/matrixLen][i%matrixLen]
	}

	goodFunc := func(i int) bool {
		return matrixToRow(i) <= target
	}

	l, r := 0, len(matrix)*matrixLen
	for r-l > 1 {
		m := (l + r) / 2
		if goodFunc(m) {
			l = m
			continue
		}

		r = m
	}

	return matrixToRow(l) == target
}
