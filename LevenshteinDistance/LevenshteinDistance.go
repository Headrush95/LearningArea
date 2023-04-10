package LevenshteinDistance

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// center вспомогательная функция для PrintMatrix, чтобы центрировать элементы
func center(a int, w int) string {
	s := strconv.Itoa(a)
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}

// PrintMatrix печатает в Stdout двумерную матрицу
func PrintMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(center(m[i][j], 3))
		}
		fmt.Println()
	}
}

// getMin возвращает наименьшее значение для данных типа int
func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// LevDist реализует алгоритм Левенштейна по поиску наименьшего количества изменений одного (comp) слова, чтобы получить второе (org)
func LevDist(org string, comp string) int {
	lOrg := utf8.RuneCountInString(org)
	lComp := utf8.RuneCountInString(comp)
	orgRune := []rune(org)
	compRune := []rune(comp)

	//Создаем начальную матрицу
	matrix := make([][]int, lOrg+1)
	for i, j := 0, 0; i <= lOrg; i++ {
		matrix[i] = make([]int, lComp+1, lComp+1)
		matrix[i][0] = j
		j++
	}

	for i := 0; i <= lComp; i++ {
		matrix[0][i] = i
	}
	for i := 1; i <= lOrg; i++ {
		for j := 1; j <= lComp; j++ {
			var diagInc = 0
			matrix[i][j] = getMin(matrix[i-1][j]+1, matrix[i][j-1]+1)
			var isSameLetter = orgRune[i-1] == compRune[j-1]
			if !isSameLetter {
				diagInc = 1
			}
			matrix[i][j] = getMin(matrix[i][j], matrix[i-1][j-1]+diagInc)
		}
	}

	PrintMatrix(matrix)
	return matrix[lOrg-1][lComp-1]
}
