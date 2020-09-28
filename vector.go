package vector

import (
	"errors"
)

type Vector []float64

var ErrVectorSize = errors.New("not compatible with vectors of different dimension")

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func (a Vector) Add(b Vector) (*Vector, error) {
	if len(a) != len(b) {
		return nil, ErrVectorSize
	}

	result := make(Vector, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	return &result, nil
}

func (a Vector) Sub(b Vector) (*Vector, error) {
	if len(a) != len(b) {
		return nil, ErrVectorSize
	}

	result := make(Vector, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] - b[i]
	}

	return &result, nil
}
