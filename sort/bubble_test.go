package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
		want  []float64
	}{
		{
			name:  "Reverse sorted",
			input: []float64{5.0, 4.0, 3.0, 2.0, 1.0},
			want:  []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]float64, len(tc.input))
			copy(input, tc.input)
			BubbleSort(input)
			if !reflect.DeepEqual(input, tc.want) {
				t.Errorf("BubbleSort (%v) = %v, want %v", tc.input, input, tc.want)
			}
		})
	}
}
