package passenger

import "fmt"

type Status int

const (
	Queued  Status = 0
	Waiting Status = 1
	Boarded Status = 2
	Arrived Status = 3
)

type Passenger struct {
	ID          int
	Status      Status
	Origin      int
	Destination int
}

func NewPassenger() *Passenger {
	return &Passenger{
		ID: 0,
		// Status is the current status of the passenger on their trip: queued, onboard, or arrived.
		Status:      0,
		Origin:      0,
		Destination: 0,
	}
}

func (p *Passenger) String() string {
	return fmt.Sprintf("Passenger %d: Origin %d, Destination %d, Status %d", p.ID, p.Origin, p.Destination, p.Status)
}
