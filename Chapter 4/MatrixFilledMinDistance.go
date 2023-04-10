package Chapter_4

// min возвращает минимальное значение из переданых
func min(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < res {
			res = arr[i]
		}
	}
	return res
}

// MatrixDistanceToCorner возвращает 2D матрицу, в которой в каждой ячейке расстояние до ближайшей стороны прямоугольника
func MatrixDistanceToCorner(a, b int) [][]int {
	matrix := make([][]int, a, a)

	for i := range matrix {
		matrix[i] = make([]int, b, b)
	}

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			matrix[i][j] = min(i, j, a-1-i, b-1-j)
		}
	}

	return matrix
}
