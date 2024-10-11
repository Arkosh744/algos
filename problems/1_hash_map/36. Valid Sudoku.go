package hashMap

// https://leetcode.com/problems/valid-sudoku/description/

func isValidSudoku(board [][]byte) bool {
	// format ["5","3",".",".","7",".",".",".","."]

	// rows
	for i := 0; i < 9; i++ {
		line := make(map[byte]struct{})
		for _, cell := range board[i] {
			if cell != '.' {
				if _, ok := line[cell]; ok {
					return false
				}
				line[cell] = struct{}{}
			}
		}
	}

	// columns
	for j := 0; j < 9; j++ {
		column := make(map[byte]struct{})
		for i := 0; i < 9; i++ {
			cell := board[i][j]
			if cell != '.' {
				if _, ok := column[cell]; ok {
					return false
				}
				column[cell] = struct{}{}
			}
		}
	}

	return validateSector(board)
}

func validateSector(board [][]byte) bool {
	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			sector := make(map[byte]struct{})

			for j := 0; j < 3; j++ {
				for i := 0; i < 3; i++ {
					cell := board[x+i][y+j]
					if cell != '.' {
						if _, ok := sector[cell]; ok {
							return false
						}
						sector[cell] = struct{}{}
					}
				}
			}
		}
	}

	return true
}

/*
Оценка по времени: О(1)
Так как у нас фикс матрица и количество операций константно

Оценка по памяти: О(1)
В не зависимости от входа выделяем одинаковое количество памяти

Описание решения.
сначала проверим по строкам и по столбцам в циклах будем добавлять в каждой линии в мапу
во время добавления проверять что нет повторений, то же самое для столбцов.
Потом пройдём по секторам и проверим каждый сектор.
*/
