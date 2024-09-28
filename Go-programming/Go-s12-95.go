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

// var power int

// typ Jedi struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	Rank      string
// }

// func foo() {
// 	name := "Luke"
// }

// func (j Jedi) ForcePush(p int) int {
// 	// ...
// 	return p
// }

// type power int

// func (p power) unleash() {
// 	// ...
// }

// type ForceUser interface {
// 	ForcePush(intVar int) int
// }

package main

import (
	"fmt"
	"math"
)

func main() {
	x := 7.05

	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.1
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.15
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.2
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.25
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.3
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.35
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.4
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.45
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.5
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.55
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.6
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.65
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.7
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.75
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.8
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.85
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.9
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 7.95
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.0
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.05
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.1
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.15
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.2
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.25
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.3
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.35
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.4
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.45
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.5
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.55
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.6
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.65
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.7
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.75
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.8
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.85
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.9
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 8.95
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	// x = 9.0
	// fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))


	x = 9.05
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.1
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.15
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.2
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.25
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.3
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.35
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.4
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.45
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))

	x = 9.5
	fmt.Printf("math.Log(%v) = %v.\n", x, math.Log(x))
}
