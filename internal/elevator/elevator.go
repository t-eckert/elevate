package elevator

import (
	"github.com/t-eckert/elevate/internal/passenger"
)

type Elevator struct {
	// Id is the unique identifier of the elevator.
	ID       int
	Floor    float64
	Velocity float64

	passengers []*passenger.Passenger
}

func NewElevator() *Elevator {
	return &Elevator{
		ID:         0,
		Floor:      0,
		Velocity:   0,
		passengers: []*passenger.Passenger{},
	}
}

func (e *Elevator) AddPassenger(p *passenger.Passenger) {
	e.passengers = append(e.passengers, p)
}

func (e *Elevator) RemovePassengerByID(id int) *passenger.Passenger {
	for i, p := range e.passengers {
		if p.ID == id {
			e.passengers = append(e.passengers[:i], e.passengers[i+1:]...)
			return p
		}
	}

	return nil
}
