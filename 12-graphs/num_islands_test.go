package graphs

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "three islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "one big island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "no land",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "single cell land",
			grid: [][]byte{{'1'}},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]byte{{'0'}},
			want: 0,
		},
		{
			name: "diagonal not connected",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			want: 5,
		},
		{
			name: "horizontal line",
			grid: [][]byte{{'1', '1', '1', '1', '1'}},
			want: 1,
		},
		{
			name: "vertical line",
			grid: [][]byte{{'1'}, {'1'}, {'1'}, {'1'}},
			want: 1,
		},
		{
			name: "checkerboard",
			grid: [][]byte{
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
				{'1', '0', '1', '0'},
			},
			want: 6,
		},
		{
			name: "empty grid",
			grid: [][]byte{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumIslands(tt.grid)
			if got != tt.want {
				t.Errorf("NumIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}
