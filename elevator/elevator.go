package elevator

import "github.com/t-eckert/elevate/passenger"

type Elevator struct {
	id string

	// Floor is the floor where the elevator is currently located.
	// It is a floating point number as the elevator may be between two floors.
	Floor float64
	// Velocity is the speed and direction of the elevator. If velocity is positive,
	// the elevator is ascending. If the velocity is negative, it is descending.
	Velocity float64
	// Passengers is a map of all passenger ids to passenger objects.
	// These passengers will all be assigned to this elevator, but may
	// have a status of queued, boarded, or arrived.
	Passengers map[string]*passenger.Passenger
}

func NewElevator(id string) *Elevator {
	return &Elevator{
		id:         id,
		Floor:      0,
		Velocity:   0,
		Passengers: make(map[string]*passenger.Passenger, 50),
	}
}

// Onboard returns the number of passengers currently on board the elevator.
func (e *Elevator) Onboard() int {
	onboard := 0
	for _, p := range e.Passengers {
		if p.Status == passenger.Boarded {
			onboard++
		}
	}
	return onboard
}
