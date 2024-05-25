package shipinfogen

import "errors"

func (s *Ship) UsedCargoSpace() uint64 {
	total_cargo := uint64(0)
	for _, cargo := range s.cargoManifest {
		total_cargo += cargo
	}
	return uint64(total_cargo)
}

func (s *Ship) LoadCargo(cargo string, units uint64) (error) {
	if units > s.UsedCargoSpace() {
		return errors.New("Attempt to load too much cargo")
	}

	return nil
}

// Check a ship's cargo.
//
// uint64 - Amount of cargo. Returns zero if cargo not found.
func (s *Ship) CheckCargo(cargo string) uint64 {
	if amount, exists := s.cargoManifest[cargo]; exists {
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
func (s *Ship) UnloadCargo(cargo string, units uint64) (uint64, error) {
	withdrawn := uint64(0)
	var err error = nil
	if (units > s.CheckCargo(cargo)) {
		withdrawn = s.cargoManifest[cargo] 
	}

	return withdrawn, err
}
