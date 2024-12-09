package list

func isPalindrome(head *ListNode) bool {
	// в задаче используется реверс и поиск средней ноды (переопределил в методе для удобства в одном файле, а не пакете)
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		curr := head

		for curr != nil {
			tmp := curr
			curr = curr.Next
			tmp.Next = prev
			prev = tmp
		}

		return prev
	}

	middle := func(head *ListNode) *ListNode {
		plus1 := head
		plus2 := head

		for plus2 != nil && plus2.Next != nil {
			plus1 = plus1.Next

			plus2 = plus2.Next.Next
		}

		return plus1
	}

	firstHalfEndNode := middle(head)
	secondHalfBeginNode := reverse(firstHalfEndNode)

	p1 := head
	p2 := secondHalfBeginNode
	// тут важно понимать как будет выглядеть связный список после манипуляций c поворотом

	//      p1            p2
	// in:  1  ->  2  ->  3
	// out: 1  ->  2  <-  3
	//            /
	//     None <|
	// т е из "2" идет в None (т.к. перевернули только 2ю половину)

	//      p1                   p2
	// in:  1  ->  2  ->  3  ->  4
	// out: 1  ->  2  <-  3  <-  4
	//                   /
	//            None <|

	// теперь очевидно, что мы должны сравнивать значения в
	// нодах p1 и p2 пока p2 не станет None
	for p2 != nil && p1 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}
