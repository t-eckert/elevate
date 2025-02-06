package elevator_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/t-eckert/elevate/internal/elevator"
)

func TestClampedVelocity(t *testing.T) {
	cases := []struct {
		current     float64
		destination float64
		expected    float64
	}{
		{0, 10, 1},
		{10, 0, -1},
		{4.6, 5, 0.4},
		{4.9, 5, 0.1},
		{5.1, 5, -0.1},
		{5, 5, 0},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%f to %f should be %f", tc.current, tc.destination, tc.expected), func(t *testing.T) {
			actual := elevator.Clamp(elevator.Velocity(tc.current, tc.destination), -1, 1)
			require.True(t, almostEqual(tc.expected, actual))
		})
	}
}

func almostEqual(a, b float64) bool {
	return math.Abs(float64(a-b)) < 0.001
}
