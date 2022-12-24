package elevator_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/t-eckert/elevate/elevator"
	"github.com/t-eckert/elevate/passenger"
)

var (
	queued  = map[string]*passenger.Passenger{}
	boarded = map[string]*passenger.Passenger{}
	arrived = map[string]*passenger.Passenger{}
)

func TestBoarded(t *testing.T) {
	e := elevator.NewElevator("test")
	expected := len(boarded)

	for id, p := range queued {
		e.Passengers[id] = p
	}
	for id, p := range boarded {
		e.Passengers[id] = p
	}
	for id, p := range arrived {
		e.Passengers[id] = p
	}

	actual := e.Onboard()

	require.Equal(t, expected, actual)
}
