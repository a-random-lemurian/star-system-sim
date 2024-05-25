package shipinfogen

import (
	"lemuria/spaceport/shipnamegen"
)

type ShipGenerator struct {
	shipsSoFar uint64
	names      shipnamegen.PhraseSet
	phraseName string
}

func (sg *ShipGenerator) GenerateName() string {
	return sg.names.GenerateString(sg.phraseName)
}

func (sg *ShipGenerator) GenerateShip() Ship {
	sg.shipsSoFar++
	return Ship{
		sig_incr_id:   sg.shipsSoFar,
		name:          sg.GenerateName(),
		cargoCapacity: 10000,
	}
}

func CreateDefaultGenerator() ShipGenerator {
	return ShipGenerator{
		shipsSoFar: 0,
		names:      shipnamegen.DefaultPhraseSet(),
		phraseName: "shipname",
	}
}
