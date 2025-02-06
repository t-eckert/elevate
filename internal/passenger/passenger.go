package passenger

type Passenger struct {
	ID int
}

func NewPassenger() *Passenger {
	return &Passenger{
		ID: 0,
	}
}
