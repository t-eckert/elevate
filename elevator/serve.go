package elevator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Serve runs the elevator continuously, moving the elevator up and down,
// picking up passengers, and dropping them off.
func Serve(ctx context.Context, e *Elevator) {
	// Register the elevator with the controllers.
	go func() {
		b, err := json.Marshal(e)
		if err != nil {
			log.Fatal(err.Error())
		}
		resp, err := http.Post(fmt.Sprintf("%s/elevators", e.config.ControllerAddress), "application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err.Error())
		}
		if resp.StatusCode != http.StatusOK {
			log.Fatal(resp.Status)
		}

		<-ctx.Done()

		// TODO actually run this on cancellation
		fmt.Println("Deregistering")
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/elevators", e.config.ControllerAddress), bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err.Error())
		}
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}
		if resp.StatusCode != http.StatusOK {
			log.Fatal(resp.Status)
		}
	}()

	// Move the elevator.
	go func() {
		for {
			e.Navigate()
			e.Move()

			select {
			case <-time.After(10 * time.Millisecond):
			case <-ctx.Done():
				return
			}
		}
	}()

	// Serve the passengers.
	go func() {
		for {
			path := e.Path
			if len(path) == 0 {
				continue
			}

			if e.AtFloor(path[0]) {
				e.PickupAndDropoff()
			}

			select {
			case <-time.After(time.Millisecond):
			case <-ctx.Done():
				return
			}
		}
	}()
}
