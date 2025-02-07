# Notes

I want to improve my understanding of the Raft consensus algorithm.
I think this could be a good project to base a talk on.
This is inspired by an interview question I got at HashiCorp.


I'm working on a demo project in Go to explore Raft and client-server architecture. The project is an elevator management system that works in a "cloud native" way. Elevators are clients that can accept passengers, navigate themselves to a given floor (simulated), and drop off passengers. The Control Plane is a server that can receive requests from passengers and route them to the correct elevator based on an "availability" score. The Control Plane can run as multiple, raft-replicated servers too.


## Components

### `controlplane`

This is a raft-replicated server that can take requests from passengers and route them to the correct elevator based on an "availability" score.

### `elevator`

This is a client that can accept passengers, navigate themselves to a given floor (simulated), and drop off passengers.

- [x] The elevator component creates a new instance of the elevator
- [ ] 
- [ ] 
- [ ] 
- [ ] 
- [ ] 
- [ ] 

How should the elevator step through its path?
I made a simple pathfinder.
Should the Elevator manage its own current location and velocity or should the server control that?


### `CLI`

There should also be a CLI for sending passengers to the Control Plane.

### `UI`

I want to write a UI in Svelte (that is served by the Go server) to show the current state of the elevators and passengers and allow the user to send passengers to the Control Plane.



