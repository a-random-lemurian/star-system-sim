package starsystem

import (
	"lemuria/spaceport/shipinfogen"
	"log"
	"math/rand"
	"sync"
	"time"
)

type StarSystem struct {
	name           string
	shipEntryPoint chan *TravelingShip
	position       Point
	connections    []*StarSystem
}

type TravelingShip struct {
	lastSystem *StarSystem
	ship       *shipinfogen.Ship
	travels    int64
}

func CreateTravelingShip(ship *shipinfogen.Ship) *TravelingShip {
	return &TravelingShip{
		lastSystem: nil,
		ship:       ship,
	}
}

func CreateStarSystem(name string, posX float64, posY float64) *StarSystem {
	return &StarSystem{
		name:           name,
		shipEntryPoint: make(chan *TravelingShip),
		position: Point{posX, posY},
	}
}

func (sys *StarSystem) randomConnected() *StarSystem {
	for {
		newSystem := sys.connections[rand.Intn(len(sys.connections))]
		if newSystem != sys {
			return newSystem
		}
	}
}

func (sys *StarSystem) sendShip(ship *TravelingShip, to *StarSystem) {
	dist := Distance(sys.position, to.position)
	log.Printf("%v: Travel time %v secs -> %v", ship.ship.String(), dist, to.name)
	time.Sleep(time.Duration(dist) * time.Second)
	to.shipEntryPoint <- ship
}

func (sys *StarSystem) receiveShip(incoming *TravelingShip) {
	if incoming.travels != 0 {
		log.Printf("%v: %v -> %v",
		incoming.ship.String(),
		incoming.lastSystem.name,
		sys.name)
	}
	incoming.travels++
}

func (sys *StarSystem) processShip(incoming *TravelingShip) {
	destination := sys.randomConnected()
	incoming.lastSystem = sys
	sys.sendShip(incoming, destination)
}

func (sys *StarSystem) Connect(in *StarSystem) {
	sys.connections = append(sys.connections, in)
}

func (sys *StarSystem) AddShip(incoming *TravelingShip) {
	incoming.lastSystem = sys
	sys.shipEntryPoint <- incoming
}

func (sys *StarSystem) Duty(wg *sync.WaitGroup) {
	defer wg.Done()
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
