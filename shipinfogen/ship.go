package shipinfogen

import (
	"fmt"
	"lemuria/spaceport/cargo"
)

type Ship struct {
	name          string
	sig_incr_id   uint64
	Cargo         cargo.Cargo
}

func (s *Ship) String() string {
	return fmt.Sprintf("%s (#%d)", s.name, s.sig_incr_id)
}
