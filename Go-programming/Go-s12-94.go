// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// type FootballPlayer struct {
// 	stamina int
// 	power   int
// }

// func (f FootballPlayer) KickBall() {
// 	shot := f.stamina + f.power

// 	fmt.Printf("I'm kicking the ball %v.\n", shot)
// }

// func main() {
// 	team := make([]FootballPlayer, 11)

// 	for i := 0; i < len(team); i++ {
// 		team[i] = FootballPlayer{
// 			stamina: rand.Intn(10),
// 			power:   rand.Intn(10),
// 		}
// 	}

// 	for i := 0; i < len(team); i++ {
// 		team[i].KickBall()
// 	}
// }

// package main

// import "fmt"

// type ninjaStar struct {
// 	Owner string
// }

// func (n ninjaStar) attack() {
// 	fmt.Printf("%v is throwing the star.\n", n.Owner)
// }

// func main() {
// 	stars := []ninjaStar{{Owner: "Wallace"}, {Owner: "Bill"}}

// 	for _, star := range stars {
// 		star.attack()
// 	}

// 	fmt.Println()
// }

package main

import "fmt"

type ninjaStar struct {
	Owner string
}

func (n ninjaStar) attack() {
	fmt.Printf("%v it throwing to the star.\n", n.Owner)
}

type ninjaSword struct {
	Owner string
}

func (n ninjaSword) attack() {
	fmt.Printf("%v swinging ninja sword.\n", n.Owner)
}

type ninjaWeapon interface {
	attack()
}

func attackFun(weapon ninjaWeapon) {
	weapon.attack()
}

func main() {
	stars := []ninjaStar{{Owner: "Wallace"}, {Owner: "Bill"}}

	for _, star := range stars {
		star.attack()
	}

	fmt.Println()

	swords := []ninjaSword{{Owner: "Will"}, {Owner: "Wallace"}}

	for _, sword := range swords {
		sword.attack()
	}

	fmt.Println()

	weapons := make([]ninjaWeapon, 4)

	weapons[0] = stars[0]
	weapons[1] = stars[1]
	weapons[2] = swords[0]
	weapons[3] = swords[1]

	for _, weapon := range weapons {
		attackFun(weapon)
	}
}
