package elevator

func Velocity(currentFloor, destination float64) float64 {
	return destination - currentFloor
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
