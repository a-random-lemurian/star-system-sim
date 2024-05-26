package cargo

import "errors"

type Cargo struct {
	cargoCapacity uint64
	cargoManifest map[string]uint64
}

// Create a cargo storage. Set the max capacity to 0 if it is unlimited.
func CreateCargoStorage(maxCapacity uint64) Cargo {
	return Cargo{
		cargoCapacity: maxCapacity,
		cargoManifest: make(map[string]uint64),
	}
}

func (c *Cargo) IsLimited() bool {
	return c.cargoCapacity == 0
}

func (c *Cargo) UsedCargoSpace() uint64 {
	total_cargo := uint64(0)
	for _, cargo := range c.cargoManifest {
		total_cargo += cargo
	}
	return uint64(total_cargo)
}

func (c *Cargo) LoadCargo(cargo string, units uint64) (error) {
	if units > c.UsedCargoSpace() {
		return errors.New("Attempt to load too much cargo")
	}

	return nil
}

// Check a ship's cargo.
//
// uint64 - Amount of cargo. Returns zero if cargo not found.
func (c *Cargo) CheckCargo(cargo string) uint64 {
	if amount, exists := c.cargoManifest[cargo]; exists {
		return amount
	}
	return 0
}

// Unload cargo from a ship.
//
// uint64 - The number of units of cargo withdrawn.
//
// error - An error. If you attempt to withdraw more cargo than what is
// on a ship, you will get one of these.
func (c *Cargo) UnloadCargo(cargo string, units uint64) (uint64, error) {
	withdrawn := uint64(0)
	var err error = nil
	if (units > c.CheckCargo(cargo)) {
		withdrawn = c.cargoManifest[cargo] 
	}

	return withdrawn, err
}
