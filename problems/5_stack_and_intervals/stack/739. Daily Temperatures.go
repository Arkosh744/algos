package stack

/*
Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

// time O(n)
// mem O(n)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	stack := make([][2]int, 0, len(temperatures))

	for i, temperature := range temperatures {
		for len(stack) != 0 && stack[len(stack)-1][1] < temperature {
			res[stack[len(stack)-1][0]] = i - stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, [2]int{i, temperature})
	}

	return res
}

/*
в стеке всегда храним убывающую последовательность
например: [[1, 20], [3, 8], [5, 1]]
[1, 20] -> в 1 день температура была 20

пока текущая температура больше чем температура в стеке
вынимаем удаляем из стека элементы и
вычисляем для них ответ
*/
