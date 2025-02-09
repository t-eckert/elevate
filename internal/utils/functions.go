package utils

import "golang.org/x/exp/constraints"

// PointerTo returns a pointer to the given value.
func PointerTo[T any](v T) *T {
	return &v
}

// Clamp clamps the value to be between min and max.
func Clamp[T constraints.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
