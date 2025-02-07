package elevator

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/t-eckert/elevate/internal/passenger"
)

type Elevator struct {
	id         int
	floor      float64
	velocity   float64
	passengers []*passenger.Passenger

	mu sync.Mutex
}

func NewElevator(id int) *Elevator {
	return &Elevator{
		id:         id,
		floor:      0,
		velocity:   0,
		passengers: []*passenger.Passenger{},
		mu:         sync.Mutex{},
	}
}

func (e *Elevator) ID() int {
	return e.id
}

func (e *Elevator) Floor() float64 {
	return e.floor
}

func (e *Elevator) Velocity() float64 {
	return e.velocity
}

func (e *Elevator) Passengers() []*passenger.Passenger {
	return e.passengers
}

func (e *Elevator) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			e.step()
		}
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

func (e *Elevator) requestedFloors() []float64 {
	floors := make([]float64, 0, len(e.passengers))
	for _, p := range e.passengers {
		floors = append(floors, float64(p.Destination))
	}
	return floors
}

func (e *Elevator) step() {
	log.Printf("Elevator %d is on floor %f", e.id, e.floor)

	e.mu.Lock()
	defer e.mu.Unlock()

	path := NewPathfinder(e.floor, e.velocity).Pathfind(e.requestedFloors())
	if len(path) == 0 {
		return
	}

	e.velocity = Clamp(Velocity(e.floor, path[0]), -1, 1)
	e.floor += e.velocity
	time.Sleep(time.Second)
}
