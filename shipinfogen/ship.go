package shipinfogen

import (
	"fmt"
)

type Ship struct {
	name        string
	sig_incr_id uint64
}

func (s Ship) String() string {
	return fmt.Sprintf("%s (#%d)", s.name, s.sig_incr_id)
}
