package list

/*
Input: head = [1,2,3,4,5]
Output: [3,4,5]

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
*/

// time: O(n)
// mem: O(1)
func middleNode(head *ListNode) *ListNode {
	plus1 := head // идет на +1 - будет в середине
	plus2 := head // идет на +2 - дойдет до конца

	for plus2 != nil && plus2.Next != nil { // едем пока плюс2 не закончатся
		plus1 = plus1.Next

		plus2 = plus2.Next.Next
	}

	return plus1
}
