package elevator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/t-eckert/elevate/elevator"
)

func TestPathfinder(t *testing.T) {
	cases := []struct {
		currentFloor float64
		velocity     float64
		floors       []float64
		expected     []float64
	}{
		{
			currentFloor: 10,
			velocity:     1,
			floors:       []float64{2, 4, 9, 11, 20, 15},
			expected:     []float64{11, 15, 20, 9, 4, 2},
		},
		{
			currentFloor: 10,
			velocity:     -1,
			floors:       []float64{2, 4, 9, 11, 20, 15},
			expected:     []float64{9, 4, 2, 11, 15, 20},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			actual := elevator.NewPathfinder(tc.currentFloor, tc.velocity).Pathfind(tc.floors)
			require.ElementsMatch(t, tc.expected, actual)
		})
	}
}
