package vector

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type test struct {
	receiver Vector
	input    Vector
	want     outputVector
}

type outputVector struct {
	Result *Vector
	Err    error
}

type outputFloat struct {
	Result float64
	Err    error
}

const margin = 1e-9

func diff(got, want outputVector, t *testing.T) {
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
		"(1,0,1)+(0,1,0)=(1,1,1)":            {receiver: Vector{1, 0, 1}, input: Vector{0, 1, 0}, want: outputVector{Result: &Vector{1, 1, 1}, Err: nil}},
		"(1,1,1)+(2,2,2)=(3,3,3)":            {receiver: Vector{1, 1, 1}, input: Vector{2, 2, 2}, want: outputVector{Result: &Vector{3, 3, 3}, Err: nil}},
		"(1,0)+(0,1)=(1,1)":                  {receiver: Vector{1, 0}, input: Vector{0, 1}, want: outputVector{Result: &Vector{1, 1}, Err: nil}},
		"(1,1,1)+(2,2,2,2)=error":            {receiver: Vector{1, 1, 1}, input: Vector{2, 2, 2, 2}, want: outputVector{Result: nil, Err: cmpopts.AnyError}},
		"(1,3,-1)+(-2,1,6)=(-1,4,5)":         {receiver: Vector{1, 3, -1}, input: Vector{-2, 1, 6}, want: outputVector{Result: &Vector{-1, 4, 5}, Err: nil}},
		"(1.3,3.5,-1)+(-2.2,1.7,6)=(-1,4,5)": {receiver: Vector{1.3, 3.5, -1}, input: Vector{-2.2, 1.7, 6}, want: outputVector{Result: &Vector{-0.9, 5.2, 5}, Err: nil}},
		"(1,1,1)+(0,0,0)=(1,1,1)":            {receiver: Vector{1, 1, 1}, input: Vector{0, 0, 0}, want: outputVector{Result: &Vector{1, 1, 1}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := outputVector{}
			got.Result, got.Err = tc.receiver.Add(tc.input)

			diff(got, tc.want, t)
		})
	}
}

func TestSub(t *testing.T) {
	tests := map[string]test{
		"(0,1,0)-(0,1,0)=(0,0,0)":                {receiver: Vector{0, 1, 0}, input: Vector{0, 1, 0}, want: outputVector{Result: &Vector{0, 0, 0}, Err: nil}},
		"(3,3,3)-(2,2,2)=(1,1,1)":                {receiver: Vector{3, 3, 3}, input: Vector{2, 2, 2}, want: outputVector{Result: &Vector{1, 1, 1}, Err: nil}},
		"(3,3)-(2,2)=(1,1)":                      {receiver: Vector{3, 3}, input: Vector{2, 2}, want: outputVector{Result: &Vector{1, 1}, Err: nil}},
		"(3,3,3)-(2,2,2,2)=error":                {receiver: Vector{3, 3, 3}, input: Vector{2, 2, 2, 2}, want: outputVector{Result: nil, Err: cmpopts.AnyError}},
		"(1,3,-1)+(-2,1,6)=(-2,1,6)":             {receiver: Vector{1, 3, -1}, input: Vector{-2, 1, 6}, want: outputVector{Result: &Vector{3, 2, -7}, Err: nil}},
		"(1.3,3.5,-1)-(-2.2,1.7,6)=(3.5,1.8,-7)": {receiver: Vector{1.3, 3.5, -1}, input: Vector{-2.2, 1.7, 6}, want: outputVector{Result: &Vector{3.5, 1.8, -7}, Err: nil}},
		"(1,1,1)-(0,0,0)=(1,1,1)":                {receiver: Vector{1, 1, 1}, input: Vector{0, 0, 0}, want: outputVector{Result: &Vector{1, 1, 1}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := outputVector{}
			got.Result, got.Err = tc.receiver.Sub(tc.input)

			diff(got, tc.want, t)
		})
	}
}

func TestCross(t *testing.T) {
	tests := map[string]test{
		"(1,-1,1)x(3,1,-2)=(1,5,4)":                        {receiver: Vector{1, -1, 1}, input: Vector{3, 1, -2}, want: outputVector{Result: &Vector{1, 5, 4}, Err: nil}},
		"(1.5,-3.2,7.1)x(33,2,-25.1)=(66.12,271.95,108.6)": {receiver: Vector{1.5, -3.2, 7.1}, input: Vector{33, 2, -25.1}, want: outputVector{Result: &Vector{66.12, 271.95, 108.6}, Err: nil}},
		"(5,5,5)x(1,0,4)=(20,-15,-5)":                      {receiver: Vector{5, 5, 5}, input: Vector{1, 0, 4}, want: outputVector{Result: &Vector{20, -15, -5}, Err: nil}},
		"(5,5)x(1,0,4)=error":                              {receiver: Vector{5, 5}, input: Vector{1, 0, 4}, want: outputVector{Result: nil, Err: cmpopts.AnyError}},
		"(5,5,1,2)x(1,0,4,2)=error":                        {receiver: Vector{5, 5, 1, 2}, input: Vector{1, 0, 4, 2}, want: outputVector{Result: nil, Err: cmpopts.AnyError}},
		"(1,1,1)x(0,0,0)=(0,0,0)":                          {receiver: Vector{1, 1, 1}, input: Vector{0, 0, 0}, want: outputVector{Result: &Vector{0, 0, 0}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := outputVector{}
			got.Result, got.Err = tc.receiver.Cross(tc.input)

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
		"(0,1,0)*(2)=(0,2,0)":       {receiver: Vector{0, 1, 0}, input: 2, want: &Vector{0, 2, 0}},
		"(3,3,3)*(2)=(6,6,6)":       {receiver: Vector{3, 3, 3}, input: 2, want: &Vector{6, 6, 6}},
		"(4,1)*(1.5)=(6,1.5)":       {receiver: Vector{4, 1}, input: 1.5, want: &Vector{6, 1, 1.5}},
		"(3,6,7)*(-3)=(-1,-18,-21)": {receiver: Vector{3, 6, 7}, input: -3, want: &Vector{-1, -18, -21}},
		"(4,1)*(0)=(0,0)":           {receiver: Vector{4, 1}, input: 0, want: &Vector{0, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.receiver.MultiByScalar(tc.input)

			cmp.Diff(got, tc.want)
		})
	}
}

func TestDot(t *testing.T) {
	tests := map[string]struct {
		receiver Vector
		input    Vector
		want     outputFloat
	}{
		"(12,20).(16,-5)=92":               {receiver: Vector{12, 20}, input: Vector{16, -5}, want: outputFloat{Result: 92., Err: nil}},
		"(1,2,5).(6,-1,2)=14":              {receiver: Vector{1, 2, 5}, input: Vector{6, -1, 2}, want: outputFloat{Result: 14, Err: nil}},
		"(-1.5,2.1,5.5).(6.2,-10,2)=-19.3": {receiver: Vector{-1.5, 2.1, 5.5}, input: Vector{6.2, -10, 2}, want: outputFloat{Result: -19.3, Err: nil}},
		"(12,20).(0,0)=0":                  {receiver: Vector{12, 20}, input: Vector{0, 0}, want: outputFloat{Result: 0, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := outputFloat{}
			got.Result, got.Err = tc.receiver.Dot(tc.input)

			diffResult := cmp.Diff(got.Result, tc.want.Result, cmpopts.EquateApprox(0, margin))
			if diffResult != "" {
				t.Fatalf(diffResult)
			}

			diffErr := cmp.Diff(got.Err, tc.want.Err, cmpopts.EquateErrors())
			if diffErr != "" {
				t.Fatalf(diffErr)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := map[string]struct {
		receiver Vector
		want     outputFloat
	}{
		"Length((0,1,0))=": {receiver: Vector{0, 1, 0}, want: outputFloat{Result: 1, Err: nil}},
		"Length((12,-5))=": {receiver: Vector{12, -5}, want: outputFloat{Result: 13, Err: nil}},
		"Length((0,0))=":   {receiver: Vector{0, 0}, want: outputFloat{Result: 0, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := outputFloat{}
			got.Result, got.Err = tc.receiver.Length()

			diffResult := cmp.Diff(got.Result, tc.want.Result, cmpopts.EquateApprox(0, margin))
			if diffResult != "" {
				t.Fatalf(diffResult)
			}

			diffErr := cmp.Diff(got.Err, tc.want.Err, cmpopts.EquateErrors())
			if diffErr != "" {
				t.Fatalf(diffErr)
			}
		})
	}
}
