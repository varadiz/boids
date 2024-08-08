package utils_test

import (
	"boids/utils"
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name      string
		a         utils.Vector
		b         utils.Vector
		want      float64
		tolerance float64
	}{
		{
			name:      "same points",
			a:         utils.Vector{X: 0, Y: 0},
			b:         utils.Vector{X: 0, Y: 0},
			want:      0,
			tolerance: 0.01,
		},
		{
			name:      "positive coordinates",
			a:         utils.Vector{X: 0, Y: 0},
			b:         utils.Vector{X: 3, Y: 4},
			want:      5,
			tolerance: 0.01,
		},
		{
			name:      "negative coordinates",
			a:         utils.Vector{X: -1, Y: -1},
			b:         utils.Vector{X: -4, Y: -5},
			want:      5,
			tolerance: 0.01,
		},
		{
			name:      "mixed coordinates",
			a:         utils.Vector{X: -1, Y: 2},
			b:         utils.Vector{X: 3, Y: -4},
			want:      7.2111,
			tolerance: 0.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Distance(tt.a, tt.b); !almostEqual(got, tt.want, tt.tolerance) {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
