package vector

import (
	"errors"
)

type Vector []float64

var errVectorSize = errors.New("not compatible with vectors of different dimension")

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func (v Vector) Add(w Vector) (*Vector, error) {
	if len(v) != len(w) {
		return nil, errVectorSize
	}

	result := make(Vector, len(v))

	for i := 0; i < len(v); i++ {
		result[i] = v[i] + w[i]
	}

	return &result, nil
}

func (v Vector) Sub(w Vector) (*Vector, error) {
	if len(v) != len(w) {
		return nil, errVectorSize
	}

	result := make(Vector, len(v))

	for i := 0; i < len(v); i++ {
		result[i] = v[i] - w[i]
	}

	return &result, nil
}

func (v Vector) MultiByScalar(s float64) *Vector {
	result := make(Vector, len(v))

	for i, val := range v {
		result[i] = val * s
	}

	return &result
}
