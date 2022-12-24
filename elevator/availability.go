package elevator

import "github.com/t-eckert/elevate/passenger"

// Weights for the components of availability.
// Higher weights increase the effect of that component on availability.
const (
	enrouteWeight       = 1.0
	pathAlignmentWeight = 1.0
	nearnessWeight      = 1.0
)

// CalculateAvailability takes an elevator and a passenger and returns a score for how
// available that elevator is to take a given passenger from their origin to their destination.
// A higher score means that the elevator is more likely to be assigned to take the passenger.
func CalculateAvailability(elevator Elevator, passenger passenger.Passenger) float64 {

	return 0
}
