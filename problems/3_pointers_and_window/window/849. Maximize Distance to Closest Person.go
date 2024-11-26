package window

/* https://leetcode.com/problems/maximize-distance-to-closest-person/

Input: seats = [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (i.e. seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
Example 2:

Input: seats = [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
Example 3:

Input: seats = [0,1]
Output: 1
*/

func maxDistToClosest(seats []int) int {
	var l, r, result int
	for l < len(seats) {
		for r+1 < len(seats) && seats[r] == seats[r+1] { // смещаем правый указатель определяя окно последовательных чисел
			r += 1
		}

		if seats[r] == 0 { // пересчитываем ответ, только если места были 0
			// если 0 прижат к стенке слева или справа, т. е.
			// 1 0 0 0
			//   l   r

			if l == 0 || r == len(seats)-1 { // мы у края, считаем по краевой формуле
				result = max(result, r-l+1)
			} else {
				result = max(result, (r-l+2)/2)
			}
		}

		// т.к. окно было закрыто (новое значение в ячейке), то смещаем оба указателя вперед на +1
		l = r + 1
		r = r + 1
	}

	return result
}
