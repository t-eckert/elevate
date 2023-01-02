package controller

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/t-eckert/elevate/elevator"
	"github.com/t-eckert/elevate/passenger"
)

type Controller struct {
	Elevators map[string]*elevator.Elevator
}

func NewController() *Controller {
	return &Controller{
		Elevators: make(map[string]*elevator.Elevator, 50),
	}
}

func (c *Controller) RegisterElevator(e *elevator.Elevator) {
	log.Infof("Registering elevator %s", e.Id)
	c.updateElevators()
	c.Elevators[e.Id] = e
}

func (c *Controller) AddPassenger(p *passenger.Passenger) {
	log.Infof("Adding passenger %s", p.Id)
}

func (c *Controller) updateElevators() {
	for id, e := range c.Elevators {
		resp, err := http.Get(e.Address)
		if err != nil {
			log.Error(err.Error())
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error(err.Error())
		}

		err = json.Unmarshal(body, &e)
		if err != nil {
			log.Error(err.Error())
		}

		c.Elevators[id] = e
	}
}
