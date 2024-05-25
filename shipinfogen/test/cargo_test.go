package shipinfogen_test

import (
	"lemuria/spaceport/shipinfogen"
	"testing"
)

func getTestShip() shipinfogen.Ship {
	gen := shipinfogen.CreateDefaultGenerator()
	return gen.GenerateShip()
}

func TestLoadCargo(t *testing.T) {
	ship := getTestShip()
	if ship.LoadCargo("testing material", 8000) == nil {
		if ship.CheckCargo("testing material") != 8000 {
			t.Error("Failed to properly load the Testing Material.")
		}
	}
}

func TestLoadCargoWithError(t *testing.T) {
	ship := getTestShip()
	if ship.LoadCargo("testing material", 12000) == nil {
		t.Error("Failed to provide an error when loading too much Testing Material.")
	}
}

func TestUnloadCargo(t *testing.T) {
	ship := getTestShip()
	if ship.LoadCargo("testing material", 8000) == nil {
		unloaded, err := ship.UnloadCargo("testing material", 4000)
		if (ship.CheckCargo("testing material") != 4000 &&
		    unloaded != 4000 && err != nil) {
			t.Error("Failed to properly unload the Testing Material.")
		}
	}
}

func TestUnloadCargoWithError(t *testing.T) {
	ship := getTestShip()
	if ship.LoadCargo("testing material", 8000) == nil {
		unloaded, err := ship.UnloadCargo("testing material", 9000)
		if unloaded != 8000 && err == nil {
			t.Error("Failed to provide an error when unloading too much Testing Material.")
		}
	}
}
