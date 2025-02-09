package elevator

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/t-eckert/elevate/internal/passenger"
	"github.com/t-eckert/elevate/internal/utils"
)

// Elevator is a representation of an elevator in the system.
type Elevator struct {
	// RequestChan is a channel for passengers to request an elevator.
	RequestChan chan *passenger.Passenger
	// DropoffChan is a channel for passengers to be dropped off at their destination.
	DropoffChan chan *passenger.Passenger

	id         int
	position   float64
	velocity   float64
	passengers []*passenger.Passenger
	requests   []*passenger.Passenger

	mu sync.Mutex
}

func NewElevator(id int) *Elevator {
	return &Elevator{
		RequestChan: make(chan *passenger.Passenger),
		DropoffChan: make(chan *passenger.Passenger),

		id:         id,
		position:   0,
		velocity:   0,
		passengers: []*passenger.Passenger{},

		mu: sync.Mutex{},
	}
}

func (e *Elevator) ID() int                            { return e.id }
func (e *Elevator) Position() float64                  { return e.position }
func (e *Elevator) Velocity() float64                  { return e.velocity }
func (e *Elevator) Passengers() []*passenger.Passenger { return e.passengers }
func (e *Elevator) Requests() []*passenger.Passenger   { return e.requests }

// Floor returns the current floor of the elevator.
// For the elevator to be "at" a floor, it must be within 0.01 of the floor.
// Otherwise, it is considered to be between floors and the function returns nil.
func (e *Elevator) Floor() *int {
	f := math.Floor(e.position)
	if math.Abs(e.position-f) < 0.01 {
		return utils.PointerTo(int(f))
	}

	c := math.Ceil(e.position)
	if math.Abs(e.position-c) < 0.01 {
		return utils.PointerTo(int(c))
	}

	return nil
}

// Serve runs the elevator until the context is cancelled.
func (e *Elevator) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case p := <-e.RequestChan:
			e.requests = append(e.requests, p)
			p.Status = passenger.Waiting
		default:
			e.move()
			e.dropoff()
			e.pickup()
		}
		time.Sleep(utils.Tick)
	}
}

func (e *Elevator) move() {
	e.mu.Lock()
	defer e.mu.Unlock()

	path := NewPathfinder(e.position, e.velocity).Pathfind(e.requestedFloors())
	if len(path) == 0 {
		return
	}

	e.velocity = Clamp(Velocity(e.position, path[0]), -0.001, 0.001)
	e.position += e.velocity
}

func (e *Elevator) dropoff() {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, p := range e.passengers {
		if e.Floor() != nil && p.Destination == *e.Floor() {
			e.removePassengerByID(p.ID)
			e.DropoffChan <- p
		}
	}
}

func (e *Elevator) pickup() {
	e.mu.Lock()
	defer e.mu.Unlock()

	toPickup := []int{}
	for i, r := range e.requests {
		if e.Floor() != nil && r.Origin == *e.Floor() {
			toPickup = append(toPickup, i)
		}
	}

	pickedUp := 0
	for _, i := range toPickup {
		p := e.requests[i-pickedUp]
		e.requests = append(e.requests[:i-pickedUp], e.requests[i-pickedUp+1:]...)
		e.passengers = append(e.passengers, p)
		p.Status = passenger.Boarded
		pickedUp++
	}
}

func (e *Elevator) removePassengerByID(id int) *passenger.Passenger {
	for i, p := range e.passengers {
		if p.ID == id {
			e.passengers = append(e.passengers[:i], e.passengers[i+1:]...)
			return p
		}
	}

	return nil
}

func (e *Elevator) requestedFloors() []float64 {
	floors := make([]float64, 0, len(e.passengers)+len(e.requests))
	for _, p := range e.passengers {
		floors = append(floors, float64(p.Destination))
	}
	for _, r := range e.requests {
		floors = append(floors, float64(r.Origin))
	}
	return floors
}

func (e *Elevator) String() string {
	return fmt.Sprintf("Elevator %d: Position %f, Velocity %f, Passengers %v, Requests %v", e.id, e.position, e.velocity, e.passengers, e.requests)
}
