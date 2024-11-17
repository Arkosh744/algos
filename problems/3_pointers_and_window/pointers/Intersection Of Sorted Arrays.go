package pointers

/**
Input 1:
A: [1 2 3 3 4 5 6]
B: [3 3 5]

Output 1: [3 3 5]

Input 2:
A: [1 2 3 3 4 5 6]
B: [3 5]

Output 2: [3 5]

[3 5 7]
[1, 7]

[7]
*/

// https://www.interviewbit.com/problems/intersection-of-sorted-arrays/

// time: O(n + m) // O(n)
// mem: O(min(n, m)) // O(n)
func intersect(A []int, B []int) []int {
	var p1, p2 int
	var res []int
	// т.к. мы знаем что оба слайса возрастающие, то задаемся такими границами по движению указателей
	for p1 < len(A) && p2 < len(B) { //
		if A[p1] > B[p2] {
			p2++
		} else if A[p1] < B[p2] {
			p1++
		} else {
			res = append(res, A[p1])
			p1++
			p2++
		}
	}

	return res
}
