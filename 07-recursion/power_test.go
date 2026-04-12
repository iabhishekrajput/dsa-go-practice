package recursion

import (
	"math"
	"testing"
)

func TestPower(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{
			name: "2^10",
			x:    2.0,
			n:    10,
			want: 1024.0,
		},
		{
			name: "2^0",
			x:    2.0,
			n:    0,
			want: 1.0,
		},
		{
			name: "5^1",
			x:    5.0,
			n:    1,
			want: 5.0,
		},
		{
			name: "2^-2",
			x:    2.0,
			n:    -2,
			want: 0.25,
		},
		{
			name: "3^3",
			x:    3.0,
			n:    3,
			want: 27.0,
		},
		{
			name: "2.1^3",
			x:    2.1,
			n:    3,
			want: 9.261,
		},
		{
			name: "0^5",
			x:    0.0,
			n:    5,
			want: 0.0,
		},
		{
			name: "1^1000",
			x:    1.0,
			n:    1000,
			want: 1.0,
		},
		{
			name: "negative base even power",
			x:    -2.0,
			n:    4,
			want: 16.0,
		},
		{
			name: "negative base odd power",
			x:    -2.0,
			n:    3,
			want: -8.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Power(tt.x, tt.n)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("Power(%v, %d) = %v, want %v", tt.x, tt.n, got, tt.want)
			}
		})
	}
}
