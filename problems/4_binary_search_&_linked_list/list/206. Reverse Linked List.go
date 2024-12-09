package list

/*
Input: head = [1,2,3,4,5y]
Output: [5,4,3,2,1]
*/

// time: O(n)
// mem: O(1)
func reverseList(head *ListNode) *ListNode {
	// на каждом шаге curr двигаем по листу на 1 ноду вперед,
	// а prev - поддерживает массив в котором будет ответ, т е

	// если есть изначально массив
	// 1 -> 2 -> 3 -> 4 -> 5 -> nil

	// то через несколько шагов он будет таким
	// nil <- 1 <- 2      3 -> 4 -> 5 -> nil
	//            prev    curr

	// таким образом в prev мы получим конечный ответ
	var prev *ListNode
	curr := head

	for curr != nil {
		tmp := curr      // сохраняем текущий ноду во временную
		curr = curr.Next // сдвигаем curr на 1 ноду вперед
		tmp.Next = prev  // меняем указатель временной ноды на prev
		prev = tmp       // меняем prev на текущий ноду
	}

	return prev // возвращаем указатель на конечную ноду
}
