// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// type Player interface {
// 	KickBall()
// }

// type CR7 struct {
// 	stamina int
// 	power   int
// 	SUI     int
// }

// type FootballPlayer struct {
// 	stamina int
// 	power   int
// }

// func (f FootballPlayer) KickBall() {
// 	shot := f.stamina + f.power

// 	fmt.Printf("I'm kicking the ball %v.\n", shot)
// }

// func (c CR7) KickBall() {
// 	shot := c.stamina + c.power * c.SUI

// 	fmt.Printf("CR7 is kicking the ball %v.\n", shot)
// }

// func main() {
// 	team := make([]Player, 11)

// 	for i := 0; i < len(team) - 1; i++ {
// 		team[i] = FootballPlayer{
// 			stamina: rand.Intn(10),
// 			power:   rand.Intn(10),
// 		}
// 	}

// 	team[len(team) - 1] = CR7{
// 		stamina: 10,
// 		power:   10,
// 		SUI:     10,
// 	}

// 	for i := 0; i < len(team); i++ {
// 		team[i].KickBall()
// 	}
// }
