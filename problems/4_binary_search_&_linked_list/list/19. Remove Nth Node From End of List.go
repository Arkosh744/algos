package list

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/*
time: O(n) - проходим весь лист
mem: O(1) - создаем фикс. количество доп нод
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	slow := dummy
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
