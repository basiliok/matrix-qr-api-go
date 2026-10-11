package matrix

import "fmt"

func Validate(matrix [][]float64, maxDimension int) error {
	if len(matrix) == 0 {
		return ErrEmptyMatrix
	}

	columnCount := len(matrix[0])
	for rowIndex, row := range matrix {
		if len(row) == 0 {
			return fmt.Errorf("%w: row %d is empty", ErrEmptyRow, rowIndex)
		}

		if len(row) != columnCount {
			return fmt.Errorf("%w: row %d has %d columns, expected %d", ErrNotRectangular, rowIndex, len(row), columnCount)
		}
	}

	rowCount := len(matrix)
	if rowCount > maxDimension || columnCount > maxDimension {
		return fmt.Errorf("%w: %dx%d is larger than %dx%d", ErrDimensionTooLarge, rowCount, columnCount, maxDimension, maxDimension)
	}

	return nil
}
