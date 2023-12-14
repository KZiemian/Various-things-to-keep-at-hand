// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	http.HandleFunc("/now", logDuration(getTime))
// 	fmt.Println("Server started at port 8000")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func logDuration(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		f(w, r)

// 		end := time.Now()

// 		fmt.Println("The request took", end.Sub(start))
// 	}
// }

// func getTime(w http.ResponseWriter, r *http.Request) {
// 	now := time.Now()

// 	_, err := fmt.Fprintf(w, "%s", now)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"fmt"
	// "math"
)

// const (
// 	c = 2.99_792_458
// 	h = 6.62_607_015
// )

func main() {
	// f_a := math.Pi * h * c / (480.0)

	// fmt.Printf("f_a = %v.\n", f_a)

	// fmt.Printf("0.18 / 1.22 = %v.\n", 0.18 / 1.22)
	// numeratorOfExp := 8 * math.Pow(math.Pi, 8.0) * math.Sqrt(math.Pi)
	// alphaToPowerMinus4 := 14.0 * 14.0 * 14.0 * 14.0


	// fmt.Printf("numeratorOfExp * alphaToPowerMinus4 = %v.\n",
	// 	numeratorOfExp * alphaToPowerMinus4)

	// fmt.Printf("134_544 * alphaToPowerMinus4 = %v.\n",
	// 	134_544.0 * alphaToPowerMinus4)

	// fmt.Printf("1.0 / 14.0 = %v.\n", 1.0 / 14.0)

	fmt.Printf("0.1 / 7.0 = %v.\n", 0.1 / 7.0)
}
