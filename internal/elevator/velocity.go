package elevator

// Velocity calculates the velocity of the elevator.
// The velocity is the difference between the current floor and the destination floor.
func Velocity(currentFloor, destination float64) float64 {
	return destination - currentFloor
}

// Clamp clamps the value to be between min and max.
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
