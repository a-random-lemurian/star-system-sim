package starsystem

import "lemuria/spaceport/shipinfogen"

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
		position:       Point{posX, posY},
	}
}
