package matrix

func copyMatrix(matrix [][]float64) [][]float64 {
	copiedMatrix := make([][]float64, len(matrix))
	for rowIndex, row := range matrix {
		copiedMatrix[rowIndex] = append([]float64(nil), row...)
	}

	return copiedMatrix
}

func newZeroMatrix(rowCount int, columnCount int) [][]float64 {
	zeroMatrix := make([][]float64, rowCount)
	for rowIndex := range zeroMatrix {
		zeroMatrix[rowIndex] = make([]float64, columnCount)
	}

	return zeroMatrix
}

func replaceNegativeZeros(matrix [][]float64) {
	for rowIndex, row := range matrix {
		for columnIndex, value := range row {
			if value == 0 {
				matrix[rowIndex][columnIndex] = 0
			}
		}
	}
}
