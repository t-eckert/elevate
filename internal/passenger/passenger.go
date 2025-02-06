package passenger

type Status int

const (
	Queued  Status = 0
	Boarded Status = 1
	Arrived Status = 2
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
