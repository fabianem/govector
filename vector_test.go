package vector

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type test struct {
	receiver Vector
	input    Vector
	want     output
}

type output struct {
	Result *Vector
	Err    error
}

func diff(got, want output, t *testing.T) {
	const margin = 1e-9

	diffResult := cmp.Diff(got.Result, want.Result, cmpopts.EquateApprox(0, margin))
	if diffResult != "" {
		t.Fatalf(diffResult)
	}

	diffErr := cmp.Diff(got.Err, want.Err, cmpopts.EquateErrors())
	if diffErr != "" {
		t.Fatalf(diffErr)
	}
}

func TestAdd(t *testing.T) {
	tests := map[string]test{
		"(1,0,1)+(0,1,0)=(1,1,1)":            {receiver: Vector{1, 0, 1}, input: Vector{0, 1, 0}, want: output{Result: &Vector{1, 1, 1}, Err: nil}},
		"(1,1,1)+(2,2,2)=(3,3,3)":            {receiver: Vector{1, 1, 1}, input: Vector{2, 2, 2}, want: output{Result: &Vector{3, 3, 3}, Err: nil}},
		"(1,0)+(0,1)=(1,1)":                  {receiver: Vector{1, 0}, input: Vector{0, 1}, want: output{Result: &Vector{1, 1}, Err: nil}},
		"(1,1,1)+(2,2,2,2)=error":            {receiver: Vector{1, 1, 1}, input: Vector{2, 2, 2, 2}, want: output{Result: nil, Err: cmpopts.AnyError}},
		"(1,3,-1)+(-2,1,6)=(-1,4,5)":         {receiver: Vector{1, 3, -1}, input: Vector{-2, 1, 6}, want: output{Result: &Vector{-1, 4, 5}, Err: nil}},
		"(1.3,3.5,-1)+(-2.2,1.7,6)=(-1,4,5)": {receiver: Vector{1.3, 3.5, -1}, input: Vector{-2.2, 1.7, 6}, want: output{Result: &Vector{-0.9, 5.2, 5}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := output{}
			got.Result, got.Err = tc.receiver.Add(tc.input)

			diff(got, tc.want, t)
		})
	}
}

func TestSub(t *testing.T) {
	tests := map[string]test{
		"(0,1,0)-(0,1,0)=(0,0,0)":                {receiver: Vector{0, 1, 0}, input: Vector{0, 1, 0}, want: output{Result: &Vector{0, 0, 0}, Err: nil}},
		"(3,3,3)-(2,2,2)=(1,1,1)":                {receiver: Vector{3, 3, 3}, input: Vector{2, 2, 2}, want: output{Result: &Vector{1, 1, 1}, Err: nil}},
		"(3,3)-(2,2)=(1,1)":                      {receiver: Vector{3, 3}, input: Vector{2, 2}, want: output{Result: &Vector{1, 1}, Err: nil}},
		"(3,3,3)-(2,2,2,2)=error":                {receiver: Vector{3, 3, 3}, input: Vector{2, 2, 2, 2}, want: output{Result: nil, Err: cmpopts.AnyError}},
		"(1,3,-1)+(-2,1,6)=(-2,1,6)":             {receiver: Vector{1, 3, -1}, input: Vector{-2, 1, 6}, want: output{Result: &Vector{3, 2, -7}, Err: nil}},
		"(1.3,3.5,-1)-(-2.2,1.7,6)=(3.5,1.8,-7)": {receiver: Vector{1.3, 3.5, -1}, input: Vector{-2.2, 1.7, 6}, want: output{Result: &Vector{3.5, 1.8, -7}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := output{}
			got.Result, got.Err = tc.receiver.Sub(tc.input)

			diff(got, tc.want, t)
		})
	}
}

func TestMultiByScalar(t *testing.T) {
	tests := map[string]struct {
		receiver Vector
		input    float64
		want     *Vector
	}{
		"(0,1,0)x(2)=(0,2,0)":       {receiver: Vector{0, 1, 0}, input: 2, want: &Vector{0, 2, 0}},
		"(3,3,3)x(2)=(6,6,6)":       {receiver: Vector{3, 3, 3}, input: 2, want: &Vector{6, 6, 6}},
		"(4,1)x(1.5)=(6,1.5)":       {receiver: Vector{4, 1}, input: 1.5, want: &Vector{6, 1, 1.5}},
		"(3,6,7)x(-3)=(-1,-18,-21)": {receiver: Vector{3, 6, 7}, input: -3, want: &Vector{-1, -18, -21}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.receiver.MultiByScalar(tc.input)

			cmp.Diff(got, tc.want)
		})
	}
}
