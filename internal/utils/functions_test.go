package utils_test

import (
	"testing"

	"github.com/t-eckert/elevate/internal/utils"
)

func TestClamp(t *testing.T) {
	// Integer tests
	t.Run("Integer Tests", func(t *testing.T) {
		tests := []struct {
			name     string
			value    int
			min      int
			max      int
			expected int
		}{
			{"within range", 5, 0, 10, 5},
			{"below min", -5, 0, 10, 0},
			{"above max", 15, 0, 10, 10},
			{"at min", 0, 0, 10, 0},
			{"at max", 10, 0, 10, 10},
			{"negative range", -5, -10, -1, -5},
			{"zero range", 5, 0, 0, 0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := utils.Clamp(tt.value, tt.min, tt.max)
				if result != tt.expected {
					t.Errorf("Clamp(%d, %d, %d) = %d; want %d",
						tt.value, tt.min, tt.max, result, tt.expected)
				}
			})
		}
	})

	// Float tests
	t.Run("Float Tests", func(t *testing.T) {
		tests := []struct {
			name     string
			value    float64
			min      float64
			max      float64
			expected float64
		}{
			{"within range", 5.5, 0.0, 10.0, 5.5},
			{"below min", -5.5, 0.0, 10.0, 0.0},
			{"above max", 15.5, 0.0, 10.0, 10.0},
			{"at min", 0.0, 0.0, 10.0, 0.0},
			{"at max", 10.0, 0.0, 10.0, 10.0},
			{"fractional values", 1.5, 1.1, 1.9, 1.5},
			{"negative range", -5.5, -10.0, -1.0, -5.5},
			{"zero range", 5.5, 0.0, 0.0, 0.0},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := utils.Clamp(tt.value, tt.min, tt.max)
				if result != tt.expected {
					t.Errorf("Clamp(%f, %f, %f) = %f; want %f",
						tt.value, tt.min, tt.max, result, tt.expected)
				}
			})
		}
	})

	// Mixed type compilation test
	t.Run("Type Parameters", func(t *testing.T) {
		// These lines should compile
		_ = utils.Clamp(int8(5), int8(0), int8(10))
		_ = utils.Clamp(int16(5), int16(0), int16(10))
		_ = utils.Clamp(int32(5), int32(0), int32(10))
		_ = utils.Clamp(int64(5), int64(0), int64(10))
		_ = utils.Clamp(float32(5.5), float32(0.0), float32(10.0))
		_ = utils.Clamp(float64(5.5), float64(0.0), float64(10.0))
	})
}

// TestClampEdgeCases tests specific edge cases that might be problematic
func TestClampEdgeCases(t *testing.T) {
	t.Run("Edge Cases", func(t *testing.T) {
		// Test with min > max
		if got := utils.Clamp(5, 10, 0); got != 10 {
			t.Errorf("Clamp with min > max: got %v, want %v", got, 10)
		}

		// Test with all equal values
		if got := utils.Clamp(5, 5, 5); got != 5 {
			t.Errorf("Clamp with all equal values: got %v, want %v", got, 5)
		}
	})
}
