package main

import (
	"lemuria/spaceport/shipinfogen"
	"lemuria/spaceport/starsystem"
	"log"
	"sync"
)

func main() {
    log.Printf("Starting...\n")

    sg := shipinfogen.CreateDefaultGenerator()

    sol := starsystem.CreateStarSystem("Sol")
    acen := starsystem.CreateStarSystem("Alpha Centauri")
    barnard := starsystem.CreateStarSystem("Barnard's Star")
    groom34 := starsystem.CreateStarSystem("Groombridge 34")

    sol.Connect(acen)
    acen.Connect(sol)
    acen.Connect(barnard)
    barnard.Connect(acen)
    acen.Connect(groom34)
    groom34.Connect(acen)

    var wg sync.WaitGroup
    wg.Add(3) // Number of star systems

    go func() {
        sol.Duty(&wg)
        wg.Done()
    }()
    go func() {
        acen.Duty(&wg)
        wg.Done()
    }()
    go func() {
        barnard.Duty(&wg)
        wg.Done()
    }()
    go func() {
        groom34.Duty(&wg)
        wg.Done()
    }()

    for i := 0; i < 10000; i++ {
		go func() {
			ship := sg.GenerateShip()
			log.Printf("Init ship %v", ship)
			ts := starsystem.CreateTravelingShip(&ship)
			acen.AddShip(ts)
		}()
    }

    wg.Wait()

    log.Printf("%v, %v, %v", sol, acen, barnard)
}
