package intervals

/*
Input: firstList = [ [0,2],[5,10],[13,23],[24,25]], secondList = [ [1,5],[8,12],[15,24],[25,26]]
Output: [ [1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Example 2:

Input: firstList = [ [1,3],[5,9]], secondList = []
Output: []
*/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	isOverlapping := func(prev []int, next []int) bool {
		return max(prev[0], next[0]) <= min(prev[1], next[1])
	}

	overlapTwoIntervals := func(i []int, j []int) []int {
		return []int{max(i[0], j[0]), min(i[1], j[1])}
	}

	var res [][]int
	var p1, p2 int
	for p1 < len(firstList) && p2 < len(secondList) {
		if isOverlapping(firstList[p1], secondList[p2]) { // если пересекаются - добавляем в ответ
			res = append(res, overlapTwoIntervals(firstList[p1], secondList[p2]))
		}

		// двигаем указатель вперед у интервала с меньшим концом
		if firstList[p1][1] < secondList[p2][1] {
			p1 += 1
		} else {
			p2 += 1
		}
	}

	return res
}
