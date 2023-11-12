// package main

// import (
// 	"math/rand"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	// Get some data to dipslay in our plot.
// 	rand.Seed(int64(0))

// 	n := 10
// 	uniform := make(plotter.Values, n)
// 	normal := make(plotter.Values, n)
// 	expon := make(plotter.Values, n)

// 	for i := 0; i < n; i++ {
// 		uniform[i] = rand.Float64()
// 		normal[i] = rand.NormFloat64()
// 		expon[i] = rand.ExpFloat64()
// 	}

// 	// Create the plot and set its title and axis label.
// 	p := plot.New()

// 	p.Title.Text = "Horizontal box plotes"
// 	p.X.Label.Text = "Values"

// 	// Make horizontal boxes for our data and add them to the plot.
// 	w := vg.Points(20)
// 	b0, err := plotter.MakeHorizBoxPlot(w, 0, uniform)

// 	if err != nil {
// 		panic(err)
// 	}

// 	b1, err := plotter.MakeHorizBoxPlot(w, 1, normal)

// 	if err != nil {
// 		panic(err)
// 	}

// 	b2, err := plotter.MakeHorizBoxPlot(w, 2, expon)

// 	if err != nil {
// 		panic(err)
// 	}

// 	p.Add(b0, b1, b2)

// 	// Set the Y axis of the plot to nominal with
// 	// the given names for y = 0, y = 1, y = 2.
// 	p.NominalY("Uniform\nDistribution", "Normal\nDistribution",
// 		"Exponential\nDistribution")

// 	if err := p.Save(4*vg.Inch, 3*vg.Inch,
// 		"GoNum-Plot-s01-08-A.png"); err != nil {

// 		panic(err)
// 	}
// }

package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Get some data to display in our plot.
	rand.Seed(int64(0))

	n := 10
	uniform := make(plotter.Values, n)
	normal := make(plotter.Values, n)
	expon := make(plotter.Values, n)

	for i := 0; i < n; i++ {
		uniform[i] = rand.Float64()
		normal[i] = rand.NormFloat64()
		expon[i] = rand.ExpFloat64()
	}

	// Create the plot and set its title and axis label.
	p := plot.New()

	p.Title.Text = "Quartile plots"
	p.Y.Label.Text = "Values"

	// Make boxes for our data and add them to the plot.
	q0, err := plotter.NewQuartPlot(0, uniform)

	if err != nil {
		panic(err)
	}

	q1, err := plotter.NewQuartPlot(1, normal)

	if err != nil {
		panic(err)
	}

	q2, err := plotter.NewQuartPlot(2, expon)

	if err != nil {
		panic(err)
	}

	p.Add(q0, q1, q2)

	// Set the X axis of the plot to nominal with
	// the given names for x = 0, x = 1 and x = 2.
	p.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	if err := p.Save(3*vg.Inch, 4*vg.Inch,
		"GoNum-Plot-s01-08.png"); err != nil {

		panic(err)
	}
}
