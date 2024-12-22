package intervals

import "sort"

/*
Example 1:

Input: intervals = [ [1,3],[2,6],[8,10],[15,18]]
Output: [ [1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [ [1,4],[4,5]]
Output: [ [1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { // O(NlogN)
		return intervals[i][0] < intervals[j][0]
	}) // start <= end по условию (и это важно)

	var result [][]int
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ { // O(n)
		interval := intervals[i]

		if isOverlapping(result[len(result)-1], interval) { // [1,4],[4,5] -> [1,5]
			result[len(result)-1] = mergeTwoIntervals(result[len(result)-1], interval)
		} else { // [1,3],[4,5] -> [1,3],[4,5]
			result = append(result, interval)
		}
	}

	return result
}

func isOverlapping(prev []int, next []int) bool {
	return max(prev[0], next[0]) <= min(prev[1], next[1])
}

func mergeTwoIntervals(prev []int, next []int) []int {
	return []int{prev[0], max(prev[1], next[1])}
}
