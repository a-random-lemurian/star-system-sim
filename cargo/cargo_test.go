package cargo_test

import (
	"lemuria/spaceport/cargo"
	"testing"
)

func makeCargoStore() cargo.Cargo {
	cargo := cargo.CreateCargoStorage(10000)
	return cargo
}

func TestLoadCargo(t *testing.T) {
	cargo := makeCargoStore()
	if cargo.LoadCargo("testing material", 8000) == nil {
		if cargo.CheckCargo("testing material") != 8000 {
			t.Error("Failed to properly load the Testing Material.")
		}
	}
}

func TestLoadCargoWithError(t *testing.T) {
	cargo := makeCargoStore()
	if cargo.LoadCargo("testing material", 12000) == nil {
		t.Error("Failed to provide an error when loading too much Testing Material.")
	}
}

func TestUnloadCargo(t *testing.T) {
	cargo := makeCargoStore()
	if cargo.LoadCargo("testing material", 8000) == nil {
		unloaded, err := cargo.UnloadCargo("testing material", 4000)
		if (cargo.CheckCargo("testing material") != 4000 &&
		    unloaded != 4000 && err != nil) {
			t.Error("Failed to properly unload the Testing Material.")
		}
	}
}

func TestUnloadCargoWithError(t *testing.T) {
	cargo := makeCargoStore()
	if cargo.LoadCargo("testing material", 8000) == nil {
		unloaded, err := cargo.UnloadCargo("testing material", 9000)
		if unloaded != 8000 && err == nil {
			t.Error("Failed to provide an error when unloading too much Testing Material.")
		}
	}
}

func TestAtomicCargoTransaction(t *testing.T) {
	cargo := makeCargoStore()
	deposits := make(map[string]uint64)
	deposits["plutonium"] = 10
	cargo.AtomicCargoTransaction(deposits, nil)
	if cargo.CheckCargo("plutonium") != 10 {
		t.Errorf("Failed to do atomic cargo transaction: got %v", cargo)
	}
}
