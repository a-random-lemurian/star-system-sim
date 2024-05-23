package main

import (
	"lemuria/spaceport/shipinfogen"
	"log"
)

func main() {
	sg := shipinfogen.CreateDefaultGenerator();
	log.Printf("%v", sg.GenerateShip())
}
