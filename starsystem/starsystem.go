package starsystem

import (
	"lemuria/spaceport/cargo"
	"lemuria/spaceport/shipinfogen"
	"log"
	"math/rand"
	"sync"
	"time"
)

// A star system.
type StarSystem struct {
	name           string
	shipEntryPoint chan *TravelingShip
	position       Point
	connections    []*StarSystem
	cargo          cargo.Cargo
	industries     []*Industry
}

// A TravelingShip is a wrapper for shipinfogen.Ship.
//
// This is used to prevent circular dependencies, as ships need star system
// specific information.
type TravelingShip struct {
	lastSystem *StarSystem
	ship       *shipinfogen.Ship
	travels    int64
}

// Pick a random connected star system.
func (sys *StarSystem) randomConnected() *StarSystem {
	for {
		newSystem := sys.connections[rand.Intn(len(sys.connections))]
		if newSystem != sys {
			return newSystem
		}
	}
}

// Bring a ship to another system.
func (sys *StarSystem) sendShip(ship *TravelingShip, to *StarSystem) {
	dist := Distance(sys.position, to.position)
	log.Printf("%v: From %v, travel for %.4f secs -> %v",
		ship.ship.String(), sys.name, dist, to.name)
	time.Sleep(time.Duration(dist) * time.Second)
	to.shipEntryPoint <- ship
}

// Print output logging a ship's entry into a star system.
func (sys *StarSystem) receiveShip(incoming *TravelingShip) {
	if incoming.travels != 0 {
		log.Printf("%v: %v -> %v",
			incoming.ship.String(),
			incoming.lastSystem.name,
			sys.name)
	} else {
		log.Printf("%v: Spawned in %v.",
			incoming.ship.String(), sys.name)
	}
	incoming.travels++
}

// Internal: Simulate the action of a ship while inside a star system.
func (sys *StarSystem) processShip(incoming *TravelingShip) {
	destination := sys.randomConnected()
	incoming.lastSystem = sys
	sys.sendShip(incoming, destination)
}

// Establish a hyperlane connection between two systems.
func (sys *StarSystem) Connect(in *StarSystem) {
	sys.connections = append(sys.connections, in)
}

// Insert a new ship in the simulation.
func (sys *StarSystem) AddShip(incoming *TravelingShip) {
	incoming.lastSystem = sys
	sys.shipEntryPoint <- incoming
}

// Start performing the duty of a star system. It will begin taking ships in
// from its shipEntryPoint channel and simulate their actions, which at the
// moment consists primarily of sending them to other star systems.
func (sys *StarSystem) Duty(wg *sync.WaitGroup) {
	defer wg.Done()
	go func() {
		if len(sys.industries) == 0 {
			return
		}
		for {
			sys.SimulateProduction()
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("%v: simulating production", sys.name)
		}
	}()

	for {
		select {
		case ship := <-sys.shipEntryPoint:
			go func(ship *TravelingShip) {
				sys.receiveShip(ship)
				time.Sleep(time.Duration(rand.Intn(300)+1) * time.Second)
				go sys.processShip(ship)
			}(ship)
		}
	}
}
