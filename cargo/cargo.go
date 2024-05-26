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
	return c.cargoCapacity != 0
}

func (c *Cargo) UsedCargoSpace() uint64 {
	total_cargo := uint64(0)
	for _, cargo := range c.cargoManifest {
		total_cargo += cargo
	}
	return uint64(total_cargo)
}

func (c *Cargo) FreeSpace() uint64 {
	return c.cargoCapacity - c.UsedCargoSpace()
}

func (c *Cargo) LoadCargo(cargo string, units uint64) error {
	if units > c.FreeSpace() && c.IsLimited() {
		return errors.New("Attempt to load too much cargo")
	}

	c.cargoManifest[cargo] += units

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
	if units > c.CheckCargo(cargo) {
		withdrawn = c.cargoManifest[cargo]
	}

	return withdrawn, err
}

type transactionType int

const (
	withdraw transactionType = iota
	deposit
)

type atomicTransaction struct {
	Cargo  string
	Amount uint64
	Type  transactionType
}

func (c *Cargo) atomicCargoRollback(changes []atomicTransaction) {
	for _, change := range changes {
		switch change.Type {
		case withdraw:
			c.LoadCargo(change.Cargo, change.Amount)
		case deposit:
			c.UnloadCargo(change.Cargo, change.Amount)
		}
	}
}

// Perform multiple cargo transactions at once, rolling it all back if one
// fails.
func (c *Cargo) AtomicCargoTransaction(deposits map[string]uint64,
	withdrawals map[string]uint64) error {
	alterations := make([]atomicTransaction, 0, len(deposits)+len(withdrawals))

	performTransaction := func(cargo string, amount uint64,
		transactionType transactionType) error {
		var err error
		switch transactionType {
		case deposit:
			err = c.LoadCargo(cargo, amount)
		case withdraw:
			_, err = c.UnloadCargo(cargo, amount)
		}
		if err != nil {
			c.atomicCargoRollback(alterations)
			return err
		}
		alterations = append(alterations, atomicTransaction{
			Cargo:  cargo,
			Amount: amount,
			Type:  transactionType,
		})
		return nil
	}

	for cargo, amount := range deposits {
		if err := performTransaction(cargo, amount, deposit); err != nil {
			return err
		}
	}

	for cargo, amount := range withdrawals {
		if err := performTransaction(cargo, amount, withdraw); err != nil {
			return err
		}
	}

	return nil
}
