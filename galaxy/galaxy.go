package galaxy

import (
	"sync"

	"github.com/a-random-lemurian/star-system-sim/starsystem"
)

type Galaxy struct {
	starSystems map[string]*starsystem.StarSystem
	wg          sync.WaitGroup
}

func CreateGalaxy() Galaxy {
	return Galaxy{
		starSystems: make(map[string]*starsystem.StarSystem),
	}
}
