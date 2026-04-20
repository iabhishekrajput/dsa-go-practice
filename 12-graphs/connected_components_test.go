package graphs

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"two components", 5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{"one component chain", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
		{"all isolated", 4, [][]int{}, 4},
		{"single node", 1, [][]int{}, 1},
		{"two nodes one edge", 2, [][]int{{0, 1}}, 1},
		{"two nodes no edge", 2, [][]int{}, 2},
		{"star graph", 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, 1},
		{"three components", 6, [][]int{{0, 1}, {2, 3}, {4, 5}}, 3},
		{"cycle of three", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 1},
		{"two cycles", 6, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 4}, {4, 5}, {5, 3}}, 2},
		{"isolated with connected", 5, [][]int{{0, 1}, {2, 3}}, 3},
		{"dense graph", 4, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountComponents(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("CountComponents(%d, %v) = %d, want %d", tt.n, tt.edges, got, tt.want)
			}
		})
	}
}
