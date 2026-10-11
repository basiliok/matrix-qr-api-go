package matrix

import "errors"

var (
	ErrEmptyMatrix       = errors.New("matrix must have at least one row")
	ErrEmptyRow          = errors.New("matrix rows must have at least one column")
	ErrNotRectangular    = errors.New("matrix rows must have the same number of columns")
	ErrDimensionTooLarge = errors.New("matrix exceeds the maximum dimension")
)
