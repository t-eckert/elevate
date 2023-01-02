package elevator

import (
	"math"

	log "github.com/sirupsen/logrus"

	"github.com/t-eckert/elevate/passenger"
)

type Elevator struct {
	// Id is the unique identifier of the elevator.
	Id string
	// Address is the address where the elevator can be reached.
	Address string
	// Floor is the floor where the elevator is currently located.
	// It is a floating point number as the elevator may be between two floors.
	Floor float64
	// Velocity is the speed and direction of the elevator. If velocity is positive,
	// the elevator is ascending. If the velocity is negative, it is descending.
	// The speed of the elevator is calculated in floors per second.
	Velocity float64
	// Passengers is a map of all passenger ids to passenger objects.
	// These passengers will all be assigned to this elevator, but may
	// have a status of queued, boarded, or arrived.
	Passengers map[string]*passenger.Passenger
	// Path is the floors that the elevator will visit in the order that they will
	// be visited.
	Path []float64

	config Config
}

func NewElevator(config Config) *Elevator {
	return &Elevator{
		Id:         config.Id,
		Address:    config.Address(),
		Floor:      0,
		Velocity:   0,
		Passengers: make(map[string]*passenger.Passenger, 50),
		Path:       []float64{},
		config:     config,
	}
}

func (e *Elevator) PickupAndDropoff() {
	for _, p := range e.Passengers {
		if e.AtFloor(float64(p.Origin)) {
			log.Infof("Picking up %s", p.Id)
			p.Status = passenger.Boarded
		}
		if e.AtFloor(float64(p.Destination)) {
			log.Infof("Dropping off %s", p.Id)
			p.Status = passenger.Arrived
		}
	}
	e.updatePath()
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

func (e *Elevator) AddPassenger(p *passenger.Passenger) {
	log.Infof("Adding passenger %s", p)
	p.Elevator = e.Id
	e.Passengers[p.Id] = p
	e.updatePath()
}

// Navigate updates the velocity of the elevator
func (e *Elevator) Navigate() {
	path := e.Path

	if len(path) == 0 {
		e.Velocity = 0
		return
	}

	e.Velocity = Clamp(Velocity(e.Floor, path[0]), -1*e.config.MaxSpeed, e.config.MaxSpeed)
}

// Move updates the current floor of the elevator by 1/100 th of its current velocity.
// For continuous movement, this method should be called 100 times per second.
func (e *Elevator) Move() {
	e.Floor += e.Velocity / 100
}

// AtFloor checks if the elevator is within range to pickup or drop off a passenger.
func (e *Elevator) AtFloor(floor float64) bool {
	return math.Abs(e.Floor-floor) < 0.3
}


// String returns a shortened segment of the elevators Id.
func (e *Elevator) String() string {
	return e.Id[:9]
}

func (e *Elevator) updatePath() {
	floors := []float64{}
	for _, p := range e.Passengers {
		if p.Status == passenger.Queued {
			floors = append(floors, float64(p.Origin))
		}
		if p.Status == passenger.Boarded {
			floors = append(floors, float64(p.Destination))
		}
	}

	e.Path = NewPathfinder(e.Floor, e.Velocity).Pathfind(floors)
}
