package matrix

import "math"

func FactorizeQR(matrix [][]float64) (orthogonalMatrix [][]float64, upperTriangularMatrix [][]float64) {
	rowCount := len(matrix)
	columnCount := len(matrix[0])
	reflectionCount := min(rowCount, columnCount)

	workingMatrix := copyMatrix(matrix)
	householderScales := make([]float64, reflectionCount)
	for columnIndex := range reflectionCount {
		householderScales[columnIndex] = applyHouseholderReflection(workingMatrix, columnIndex)
	}

	upperTriangularMatrix = extractUpperTriangularMatrix(workingMatrix, reflectionCount)
	orthogonalMatrix = buildOrthogonalMatrix(workingMatrix, householderScales)

	normalizeSigns(orthogonalMatrix, upperTriangularMatrix)
	replaceNegativeZeros(orthogonalMatrix)
	replaceNegativeZeros(upperTriangularMatrix)

	return orthogonalMatrix, upperTriangularMatrix
}

func applyHouseholderReflection(workingMatrix [][]float64, columnIndex int) float64 {
	rowCount := len(workingMatrix)
	columnCount := len(workingMatrix[0])

	squaredNormBelowDiagonal := 0.0
	for rowIndex := columnIndex + 1; rowIndex < rowCount; rowIndex++ {
		value := workingMatrix[rowIndex][columnIndex]
		squaredNormBelowDiagonal += value * value
	}

	if squaredNormBelowDiagonal == 0 {
		return 0
	}

	diagonalValue := workingMatrix[columnIndex][columnIndex]
	columnNorm := math.Sqrt(diagonalValue*diagonalValue + squaredNormBelowDiagonal)
	reflectedDiagonalValue := -math.Copysign(columnNorm, diagonalValue)
	householderScale := (reflectedDiagonalValue - diagonalValue) / reflectedDiagonalValue

	householderVectorDivisor := diagonalValue - reflectedDiagonalValue
	for rowIndex := columnIndex + 1; rowIndex < rowCount; rowIndex++ {
		workingMatrix[rowIndex][columnIndex] /= householderVectorDivisor
	}
	workingMatrix[columnIndex][columnIndex] = reflectedDiagonalValue

	for targetColumnIndex := columnIndex + 1; targetColumnIndex < columnCount; targetColumnIndex++ {
		applyReflectionToColumn(workingMatrix, targetColumnIndex, workingMatrix, columnIndex, householderScale)
	}

	return householderScale
}

func applyReflectionToColumn(
	targetMatrix [][]float64,
	targetColumnIndex int,
	matrixWithHouseholderVectors [][]float64,
	reflectionIndex int,
	householderScale float64,
) {
	rowCount := len(targetMatrix)

	projection := targetMatrix[reflectionIndex][targetColumnIndex]
	for rowIndex := reflectionIndex + 1; rowIndex < rowCount; rowIndex++ {
		projection += matrixWithHouseholderVectors[rowIndex][reflectionIndex] * targetMatrix[rowIndex][targetColumnIndex]
	}
	projection *= householderScale

	targetMatrix[reflectionIndex][targetColumnIndex] -= projection
	for rowIndex := reflectionIndex + 1; rowIndex < rowCount; rowIndex++ {
		targetMatrix[rowIndex][targetColumnIndex] -= projection * matrixWithHouseholderVectors[rowIndex][reflectionIndex]
	}
}

func extractUpperTriangularMatrix(workingMatrix [][]float64, reflectionCount int) [][]float64 {
	columnCount := len(workingMatrix[0])

	upperTriangularMatrix := newZeroMatrix(reflectionCount, columnCount)
	for rowIndex := range reflectionCount {
		for columnIndex := rowIndex; columnIndex < columnCount; columnIndex++ {
			upperTriangularMatrix[rowIndex][columnIndex] = workingMatrix[rowIndex][columnIndex]
		}
	}

	return upperTriangularMatrix
}

func buildOrthogonalMatrix(workingMatrix [][]float64, householderScales []float64) [][]float64 {
	rowCount := len(workingMatrix)
	reflectionCount := len(householderScales)

	orthogonalMatrix := newZeroMatrix(rowCount, reflectionCount)
	for diagonalIndex := range reflectionCount {
		orthogonalMatrix[diagonalIndex][diagonalIndex] = 1
	}

	for reflectionIndex := reflectionCount - 1; reflectionIndex >= 0; reflectionIndex-- {
		householderScale := householderScales[reflectionIndex]
		if householderScale == 0 {
			continue
		}

		for columnIndex := reflectionIndex; columnIndex < reflectionCount; columnIndex++ {
			applyReflectionToColumn(orthogonalMatrix, columnIndex, workingMatrix, reflectionIndex, householderScale)
		}
	}

	return orthogonalMatrix
}

func normalizeSigns(orthogonalMatrix [][]float64, upperTriangularMatrix [][]float64) {
	for diagonalIndex := range upperTriangularMatrix {
		if upperTriangularMatrix[diagonalIndex][diagonalIndex] >= 0 {
			continue
		}

		for columnIndex := range upperTriangularMatrix[diagonalIndex] {
			upperTriangularMatrix[diagonalIndex][columnIndex] = -upperTriangularMatrix[diagonalIndex][columnIndex]
		}

		for rowIndex := range orthogonalMatrix {
			orthogonalMatrix[rowIndex][diagonalIndex] = -orthogonalMatrix[rowIndex][diagonalIndex]
		}
	}
}
