package coords

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name   string
		c1, c2 Coord2D
		want   float64
	}{
		{
			name: "dist 0 | 1,1 1,1",
			c1:   Coord2D{0, 0},
			c2:   Coord2D{0, 0},
			want: float64(0),
		},
		{
			name: "dist 1 | 1,1 1,2",
			c1:   Coord2D{1, 1},
			c2:   Coord2D{1, 2},
			want: float64(1),
		},
		{
			name: "dist 11,66 | -2,3 4,-7",
			c1:   Coord2D{-2, 3},
			c2:   Coord2D{4, -7},
			want: math.Float64frombits(4622754686249968546),
		},
		{
			name: "dist 1197,33 | 200,-344 -991,-221",
			c1:   Coord2D{200, -344},
			c2:   Coord2D{-991, -221},
			want: math.Float64frombits(4652980748441379786),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Distance2D(test.c1, test.c2)
			if got != test.want {
				t.Errorf("wrong result\ninput: %q\ngot:   %.10f\nwant:  %.10f",
					test.name, got, test.want)
			}
		})
	}
}
