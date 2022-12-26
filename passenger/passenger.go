package passenger

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	Queued  Status = 0
	Boarded Status = 1
	Arrived Status = 2
)

// Passenger is a person that requests a ride from their origin floor to their
// destination floor. They will be assigned to an elevator and served.
type Passenger struct {
	Id string `json:"id"`

	// Status is the current status of the passenger on their trip: queued, onboard, or arrived.
	Status Status
	// Elevator is the Id of the elevator to which the passenger is assigned.
	Elevator string
	// Origin is the floor the passenger starts their journey on.
	Origin int
	// Destination is the floor the passenger wishes to travel to.
	Destination int
}

// NewPassenger creates a new instance of a Passenger with the given id, orign,
// and destination. The new passenger will not be assigned to an elevator and
// will have a status of Queued.
func NewPassenger(id string, origin, destination int) *Passenger {
	return &Passenger{
		Id:          id,
		Status:      Queued,
		Elevator:    "",
		Origin:      origin,
		Destination: destination,
	}
}

// NewRandomPassenger creates a new instance of a passenger with a randomly
// assigned id, origin, and destination.
func NewRandomPassenger() *Passenger {
	return NewPassenger(uuid.NewString(), rand.Intn(30), rand.Intn(30))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
