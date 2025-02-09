package elevator

import "math"

// Pathfinder calculates the most efficient path to serve all of the floors
// the elevator needs to serve.
type Pathfinder struct {
	currentFloor float64
	velocity     float64
}

// NewPathfinder creates a new Pathfinder with the given current floor and velocity.
func NewPathfinder(currentFloor, velocity float64) *Pathfinder {
	return &Pathfinder{
		currentFloor: currentFloor,
		velocity:     velocity,
	}
}

// Pathfind calculates the most efficient route using a merge sort algorithm
// with a bias set by the bias method on the Pathfinder.
func (p Pathfinder) Pathfind(floors []float64) []float64 {
	if len(floors) < 2 {
		return floors
	}
	firstHalf, secondHalf := floors[:len(floors)/2], floors[len(floors)/2:]
	return p.merge(p.Pathfind(firstHalf), p.Pathfind(secondHalf))
}

func (p Pathfinder) merge(a, b []float64) []float64 {
	floors := make([]float64, 0, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if p.bias(a[i]) < p.bias(b[j]) {
			floors = append(floors, a[i])
			i++
		} else {
			floors = append(floors, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		floors = append(floors, a[i])
	}
	for ; j < len(b); j++ {
		floors = append(floors, b[j])
	}

	return floors
}

// bias returns a value that is used to sort the floors in the path. The lower
// the number, the earlier in the path it should go.
func (p Pathfinder) bias(floor float64) float64 {
	return (floor-p.currentFloor)*p.velocity +
		((floor-p.currentFloor)/math.Abs(floor-p.currentFloor))*-100
}
