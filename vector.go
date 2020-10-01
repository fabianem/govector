package vector

import (
	"errors"
	"math"
)

type Vector []float64

var errVectorsDimension = errors.New("vectors have incompatible dimensions")
var errCross3D = errors.New("the cross product of two vectors a and b is defined only in three-dimensional space")

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func (v Vector) Add(w Vector) (*Vector, error) {
	if len(v) != len(w) {
		return nil, errVectorsDimension
	}

	result := make(Vector, len(v))

	for i := 0; i < len(v); i++ {
		result[i] = v[i] + w[i]
	}

	return &result, nil
}

func (v Vector) Sub(w Vector) (*Vector, error) {
	if len(v) != len(w) {
		return nil, errVectorsDimension
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

func (v Vector) Dot(w Vector) (float64, error) {
	if len(v) != len(w) {
		return 0, errVectorsDimension
	}

	var result float64

	for i := 0; i < len(v); i++ {
		result += v[i] * w[i]
	}

	return result, nil
}

func (v Vector) Cross(w Vector) (*Vector, error) {
	if len(v) != 3 || len(w) != 3 {
		return nil, errCross3D
	}

	result := make(Vector, len(v))

	result[0] = v[1]*w[2] - v[2]*w[1]
	result[1] = v[2]*w[0] - v[0]*w[2]
	result[2] = v[0]*w[1] - v[1]*w[0]

	return &result, nil
}

func (v Vector) Length() (float64, error) {
	dot, err := v.Dot(v)
	if err != nil {
		return -1, err
	}

	return math.Sqrt(dot), nil
}
