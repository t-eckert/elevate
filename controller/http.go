package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/t-eckert/elevate/elevator"
	"github.com/t-eckert/elevate/passenger"
)

func NewIndexHandler(c *Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controller, err := json.Marshal(c)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Add("Content-Type", "application/json")
			if _, err := w.Write(controller); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func NewElevatorHandler(c *Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			var elevator elevator.Elevator
			err = json.Unmarshal(body, &elevator)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			c.RegisterElevator(&elevator)

			w.WriteHeader(http.StatusOK)
			return
		case "GET":
			controller, err := json.Marshal(c.Elevators)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Add("Content-Type", "application/json")
			if _, err := w.Write(controller); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func NewPassengerHandler(c *Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			var passenger passenger.Passenger
			err = json.Unmarshal(body, &passenger)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			c.AddPassenger(&passenger)

			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
