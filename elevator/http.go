package elevator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/t-eckert/elevate/passenger"
)

func NewIndexHandler(e *Elevator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			elevator, err := json.Marshal(e)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(elevator))

			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

}

func NewPassengerHandler(e *Elevator) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			passengers, err := json.Marshal(e.Passengers)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(passengers))

			return
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

			e.AddPassenger(&passenger)

			w.WriteHeader(http.StatusOK)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
