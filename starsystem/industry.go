package starsystem

import "log"

type Industry struct {
	industryName string
	inputs       map[string]uint64
	outputs      map[string]uint64
	id           uint64
}

var industriesSoFar = uint64(0)

func CreateIndustry(name string, inputs map[string]uint64,
	outputs map[string]uint64) Industry {
	industriesSoFar++
	return Industry{
		industryName: name,
		inputs:       inputs,
		outputs:      outputs,
		id:           industriesSoFar,
	}
}

// Add an industry to a StarSystem.
func (sys *StarSystem) AddIndustry(ind Industry) {
	sys.industries = append(sys.industries, &ind)
}

func (sys *StarSystem) SimulateProduction() {
	for _, industry := range sys.industries {
		if err := sys.cargo.AtomicCargoTransaction(
			industry.outputs, industry.inputs); err != nil {
			log.Printf("Error: %v", err)
			log.Printf("Warning: Not enough cargo in system!")
		}
	}
}
