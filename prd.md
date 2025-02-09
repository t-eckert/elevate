# PRD: Elevate - A Cloud-Native Elevator Management System

## Executive Summary

Elevate demonstrates the value of cloud-native development techniques through the example of an elevator control system. The elevators are simulated processes run by servers. They can receive passenger requests over gRPC. These passengers have an origin and a destination floor. Once the elevator receives the request, it navigates to the origin floor, picks up the passenger, then navigates to the requested floor of the passenger, and drops that passenger off by sending a message over gRPC to a configured endpoint. The path that the elevator takes as it navigates the floors is determined by an algorithm which takes into account the origin floor of every passenger waiting for the elevator and the destination floor of every boarded passenger.

Multiple elevator servers will independently. They are controlled by a control plane which exposes an API in both gRPC and as HTTP rest endpoints. This control plane consists of `n + 1` replicated servers which form a Raft quorum. A request to serve a passenger that is sent to one of the servers is forwarded to the leader of the Raft group. The leader replicates the passenger request across the Raft group. The control plane uses an availability function to determine which of the elevators it knows about should be assigned the passenger request. The passenger is sent to the elevator over gRPC.

This project aims to provide a simple example of a control plane and data plane architecture. This example will be the basis of educational content. This may include a conference talk, short book, or video.

## Problem Statement

The control plane and data plane architecture is difficult to conceptualize without concrete examples. While many implementations exist, they address a sufficiently complex business use case so as to complicate understanding the architecture. An elevator control system is complex enough to warrant an implementation of the architecture, but common enough to be understood by most engineers. 

This example implementation and the educational materials that come from it will help engineers to understand Raft replication, control plane and data plane architectures, and how to write resilient systems in Go.

## Product Overview

### Product Description

Elevate is a cloud-native elevator management system. It consists of a control plane and multiple elevator servers. The control plane is a Raft quorum that manages the communication between the external systems and the elevator servers. The elevator servers simulate the behavior of an elevator. They can receive passenger requests over gRPC and navigate to the requested floors.

### Target Users

- Learner
  - Software engineers who want to learn about cloud-native development techniques 
  - These are early to mid level engineers who are looking to learn about cloud-native development techniques. They may be looking to move into a cloud-native role or to improve their skills in their current role.
  - Software engineers who want to learn about control plane and data plane architectures
  - They are looking for a simple example of a control plane and data plane architecture that they can understand and implement themselves.
- Lecturers and Educators
  - They want to demonstrate cloud-native development techniques
  - These are educators who are looking for a simple example of a control plane and data plane architecture that they can use to teach their students.

### User Stories

| As a... | I want to... | So that... | Priority |
|---------|-------------|-------------|-----------|
| Learner | have a better understanding of cloud native architectures | to get a new job or learn a new skill | H |
| Lecturer or Educator | give a demo and teach about cloud native architectures | to communicate an important concept in backend design | M |

## Requirements

### Functional Requirements

#### Core Features

1. Elevator Simulation
   - Description: a simulation of the behavior of an elevator.
   - Acceptance Criteria: the elevator can be run, accept requests, navigate itself to floors based on the passengers waiting for it and those onboard.
   - Dependencies: None
   - Priority: 1

1. Elevator Server
   - Description: runs an elevator simulation and manages its communication with external systems.
   - Acceptance Criteria: the elevator server creates a new instance of an elevator simulation, runs that simulation, exposes gRPC endpoints for accepting passenger requests, and sends messages to the control plane when a passenger is dropped off.
   - Dependencies: Elevator Simulation
   - Priority: 2

1. Raft Implementation
   - Description: a simple implementation of the Raft consensus algorithm.
   - Acceptance Criteria: the Raft implementation can perform elections, replicate logs, and handle leader changes.
   - Dependencies: None
   - Priority: 3

1. Control Plane
   - Description: a server that manages the communication between the elevator servers and the external systems.
   - Acceptance Criteria: the control plane can receive passenger requests from the external systems, forward them to the appropriate elevator server, and manage the communication between the elevator servers and the external systems.
   - Dependencies: Elevator Server, Raft Implementation
   - Priority: 4

#### Additional Features
1. Local Deployment
   - Description: a simple deployment of the system on a local machine.
   - Acceptance Criteria: the system can be run on a local machine without any external dependencies.
   - Dependencies: Control Plane, Elevator Server
   - Priority: 5

1. Kubernetes Deployment
   - Description: a deployment of the system on Kubernetes.
   - Acceptance Criteria: the system can be run on Kubernetes using deployments and replicasets. The system should be able to handle failures and recover gracefully.
   - Dependencies: Control Plane, Elevator Server
   - Priority: 6

1. Alternative Cloud Native Deployment
   - Description: a deployment of the system on a cloud-native platform other than Kubernetes.
   - Acceptance Criteria: the system can be run on a platform other than Kubernetes, such as Nomad or AWS ECS.
   - Dependencies: Control Plane, Elevator Server
   - Priority: 7

1. User Interface
   - Description: a user interface for viewing the live state of the whole system and making passenger requests.
   - Acceptance Criteria: a user can view the state of all elevators and make requests for passengers to the system.
   - Dependencies: Control Plane, Elevator Server
   - Priority: 8

1. CLI and TUI
   - Description: a way of controlling and viewing system state from the terminal.
   - Acceptance Criteria: a user can view the state of all elevators and make requests for passengers to the system from the terminal.
   - Dependencies: Control Plane, Elevator Server
   - Priority: 9

1. Educational Materials
  - Description: a set of educational materials that explain the system and how it works.
  - Acceptance Criteria: the educational materials should explain the system, how it works, and how to deploy it.
  - Dependencies: Control Plane, Elevator Server
  - Priority: 10

1. Conference Talk
  - Description: a conference talk that explains the system and how it works.
  - Acceptance Criteria: the conference talk should explain the system, how it works, and how to deploy it.
  - Dependencies: Control Plane, Elevator Server
  - Priority: 11

1. Short Book
  - Description: a short book that explains the system and how it works.
  - Acceptance Criteria: the short book should explain the system, how it works, and how to deploy it.
  - Dependencies: Control Plane, Elevator Server
  - Priority: 12

1. Video
  - Description: a video that explains the system and how it works.
  - Acceptance Criteria: the video should explain the system, how it works, and how to deploy it.
  - Dependencies: Control Plane, Elevator Server   
  - Priority: 13

1. Disaster Recovery
  - Description: advanced disaster recovery for the system.
  - Acceptance Criteria: if an elevator fails catastrophically, the control plane should notice this and recreate the elevator from the last known snapshot.
  - Dependencies: Control Plane, Elevator Server
  - Priority: 14

1. Auto-scaling
  - Description: auto-scaling of the system.
  - Acceptance Criteria: the system should be able to scale up and down based on the number of passenger requests.
  - Dependencies: Control Plane, Elevator Server
  - Priority: 15

1. Load Testing
  - Description: load testing of the system.
  - Acceptance Criteria: the system should be able to handle a large number of passenger requests per second.
  - Dependencies: Control Plane, Elevator Server
  - Priority: 16

### Non-Functional Requirements

#### Performance

- The elevators should travel at a gated speed that would make sense for a human passenger.
- The control plane should be able to handle a large number of passenger requests per second.
- The system should be able to handle a large number of elevator servers.
- The control plane should self-recover from failures.
- The elevator servers should self-recover from failures.

#### Security

- Authentication requirements
- Data protection needs
- Compliance requirements

#### Usability

- Accessibility standards
- UI/UX requirements
- Supported devices/browsers

## Technical Specifications

### System Architecture

Two clients are available to interact with the system: a web-based user interface and a terminal-based CLI/TUI. These clients interact with the control plane either over gRPC or HTTP. The control plane is a set of `n + 1` servers behind a load balancer. The control plane can accept requests for passengers. It also stores snapshots of the state of the elevators and the passenger requests. The control plane uses a Raft implementation to replicate the state across the servers. The control plane uses an availability function to determine which elevator server should be assigned a passenger request. The control plane forwards the passenger request to the elevator server over gRPC.

### Integration Points

External system integrations are a web-based user interface and a terminal-based CLI/TUI. These systems interact with the control plane over gRPC or HTTP. The control plane interacts with the elevator servers over gRPC.

### APIs

#### Control Plane API

gRPC API:

- `ServePassenger`
  - Description: sends a passenger request to the control plane
  - Request: `PassengerRequest`
  - Response: `PassengerResponse`
- `GetElevatorState`
  - Description: gets the state of all elevators
  - Request: None
  - Response: `ElevatorState`

HTTP REST API:

- `POST /passenger`
  - Description: sends a passenger request to the control plane
  - Request: `PassengerRequest`
  - Response: `PassengerResponse`
- `GET /elevators`
  - Description: gets the state of all elevators
  - Request: None
  - Response: `ElevatorState`

#### Elevator Server API

gRPC API:

- `ServePassenger`
  - Description: sends a message to the control plane that a passenger has been served
  - Request: `PassengerRequest`
  - Response: `PassengerResponse`
- `PickupPassenger`
  - Description: sends a message to the control plane that a passenger has been picked up
  - Request: `PassengerRequest`
  - Response: `PassengerResponse`
- `DropoffPassenger`
  - Description: sends a message to the control plane that a passenger has been dropped off
  - Request: `PassengerRequest`
  - Response: `PassengerResponse`
- `GetElevatorState`
  - Description: gets the state of the elevator
  - Request: None
  - Response: `ElevatorState`

### Data Requirements

Passenger:
- ID
- Origin floor
- Destination floor
- State: queued, waiting, boarded, arrived

Elevator:
- ID
- Current floor
- Velocity
- Passengers: list of passengers
- Requests: list of passenger requests

## Design and User Experience

### Web-based User Interface

#### Key Screens and Workflows

#### Design Principles

#### User Flows

[Include key user journey maps or flow diagrams]

### Terminal-based User Interface and CLI

#### Key Screens and Workflows

#### Design Principles

#### User Flows

[Include key user journey maps or flow diagrams]

## Release Planning

### Milestones

| Milestone | Description | Target Date |
|-----------|-------------|-------------|
| [Phase 1] | [Description] | [Date] |

### Success Metrics

- [KPI 1]
  - Target:
  - Measurement method:
- [KPI 2]
  - Target:
  - Measurement method:

## Implementation Considerations

### Dependencies

- [Internal dependencies]
- [External dependencies]
- [Third-party services]

### Constraints

- [Technical constraints]
- [Business constraints]
- [Resource constraints]

### Risks and Mitigation

| Risk | Impact | Probability | Mitigation Strategy |
|------|---------|------------|---------------------|
| [Risk 1] | [H/M/L] | [H/M/L] | [Strategy] |

## Appendix

### Related Documents

- [Link to design documents]
- [Link to research findings]
- [Link to competitive analysis]

### Glossary

| Term | Definition |
|------|------------|
| [Term 1] | [Definition] |

### Open Questions

- [List any unresolved questions or decisions]
