package intervals

import "sort"

/*
Example 1:

Input: points = [ [10,16],[2,8],[1,6],[7,12]]
Output: 2
Explanation: The balloons can be burst by 2 arrows:
- Shoot an arrow at x = 6, bursting the balloons [2,8] and [1,6].
- Shoot an arrow at x = 11, bursting the balloons [10,16] and [7,12].
Example 2:

Input: points = [ [1,2],[3,4],[5,6],[7,8]]
Output: 4
Explanation: One arrow needs to be shot for each balloon for a total of 4 arrows.
Example 3:

Input: points = [ [1,2],[2,3],[3,4],[4,5]]
Output: 2
Explanation: The balloons can be burst by 2 arrows:
- Shoot an arrow at x = 2, bursting the balloons [1,2] and [2,3].
- Shoot an arrow at x = 4, bursting the balloons [3,4] and [4,5].

*/

// time: O(n log n)
// mem: O(log n) (because of sort)

func findMinArrowShots(points [][]int) int {
	isOverlapping := func(a []int, b []int) bool { // проверяет пересекутся ли интервалы
		return a[0] <= b[1] && b[0] <= a[1]
	}

	overlapTwoIntervals := func(a []int, b []int) []int {
		return []int{max(a[0], b[0]), min(a[1], b[1])}
	}

	sort.Slice(points, func(i, j int) bool { // O(n * log n)
		return points[i][0] < points[j][0]
	})

	resArrows := 1
	lastPoint := points[0]
	for _, point := range points {
		// если интервалы пересекаются, то для них используем 1 стрелу
		// и стараемся набрать как можно больше интервалов под 1 стрелу
		if isOverlapping(lastPoint, point) {
			lastPoint = overlapTwoIntervals(lastPoint, point)
			continue
		}

		// если нет пересечений значит нужна еще 1 доп стрела и обновляем
		// последний интервал (last_point)
		lastPoint = point
		resArrows += 1
	}

	return resArrows
}
