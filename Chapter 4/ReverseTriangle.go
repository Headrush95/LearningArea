package Chapter_4

// ReverseTriangle делает из нижнего треугольного массива верхний. Т.е. из массива со всеми значениями в левом нижнем углу делает массива со значениями в правом верхнем
func ReverseTriangle(matrix *[][]int) {
	lineLength := len((*matrix)[0])
	rowLength := len(*matrix)
	for i, _ := range *matrix {
		for j, _ := range (*matrix)[i] {
			(*matrix)[i][lineLength-1-j], (*matrix)[rowLength-1-i][j] = (*matrix)[rowLength-1-i][j], (*matrix)[i][lineLength-1-j]
		}

		if i > rowLength/2-2 {
			break
		}
	}

	if rowLength%2 != 0 {
		for i := 0; i < lineLength/2; i++ {
			(*matrix)[rowLength/2][i], (*matrix)[rowLength/2][lineLength-1-i] = (*matrix)[rowLength/2][lineLength-1-i], (*matrix)[rowLength/2][i]
		}
	}
}
