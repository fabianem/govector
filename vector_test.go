package vector

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestAdd(t *testing.T) {
	tests := map[string]test{
		"(1,0,1)+(0,1,0)=(1,1,1)": {receiver: Vector{1, 0, 1}, input: Vector{0, 1, 0}, want: output{Result: &Vector{1, 1, 1}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output := output{}
			output.Result, output.Err = tc.receiver.Add(tc.input)
			diff := cmp.Diff(output, tc.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := map[string]test{
		"(0,1,0)-(0,1,0)=(0,0,0)": {receiver: Vector{0, 1, 0}, input: Vector{0, 1, 0}, want: output{Result: &Vector{0, 0, 0}, Err: nil}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output := output{}
			output.Result, output.Err = tc.receiver.Sub(tc.input)
			diff := cmp.Diff(output, tc.want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
