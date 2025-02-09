package elevator

import "github.com/t-eckert/elevate/internal/passenger"

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
func CalculateAvailability(elevator *Elevator, passenger *passenger.Passenger) float64 {
	// Enroute: 1 if the elevator is enroute to the passenger's origin, 0 otherwise.
	enroute := 0.0
	// The elevator is below the passenger's origin and moving up.
	if elevator.Position() < float64(passenger.Origin) && elevator.Velocity() > 0 {
		enroute = 1.0
	}
	// The elevator is above the passenger's origin and moving down.
	if elevator.Position() > float64(passenger.Origin) && elevator.Velocity() < 0 {
		enroute = 1.0
	}

	// Path alignment: 1 if the elevator is moving in the same direction as the passenger, 0 otherwise.
	pathAlignment := 0.0
	// The elevator is moving up and the passenger is going up.
	if elevator.Velocity() > 0 && passenger.Destination > passenger.Origin {
		pathAlignment = 1.0
	}
	// The elevator is moving down and the passenger is going down.
	if elevator.Velocity() < 0 && passenger.Destination < passenger.Origin {
		pathAlignment = 1.0
	}

	// Nearness: higher if the elevator is close to the passenger's origin.
	nearness := 1.0 - (1.0 / (1.0 + 0.1*(elevator.Position()-float64(passenger.Origin))*(elevator.Position()-float64(passenger.Origin))))

	return enroute*enrouteWeight + pathAlignment*pathAlignmentWeight + nearness*nearnessWeight
}
