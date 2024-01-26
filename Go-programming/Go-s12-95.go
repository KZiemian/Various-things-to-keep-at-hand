package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

func (f FootballPlayer) KickBall() {
	shot := f.stamina + f.power

	fmt.Printf("I'm kicking the ball %v.\n", shot)
}

func main() {
	team := make([]Player, 11)

	for i := 0; i < len(team); i++ {
		team[i] = FootbalPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}

	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
