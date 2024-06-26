package shipinfogen

import (
	"github.com/a-random-lemurian/star-system-sim/shipnamegen"
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
		sig_incr_id: sg.shipsSoFar,
		name:        sg.GenerateName(),
	}
}

func CreateDefaultGenerator() ShipGenerator {
	return ShipGenerator{
		shipsSoFar: 0,
		names:      shipnamegen.DefaultPhraseSet(),
		phraseName: "shipname",
	}
}
