// package main

// import (
// 	"image/color"
// 	"math/rand"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	// Get some data to diplay in our plot.
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

// 	p.Title.Text = "Box plots"
// 	p.Y.Label.Text = "Values"

// 	// Make boxes for out data and add them to the plot.
// 	w := vg.Points(20)
// 	b0, err := plotter.NewBoxPlot(w, 0, uniform)

// 	b0.FillColor = color.RGBA{127, 188, 165, 1}

// 	if err != nil {
// 		panic(err)
// 	}

// 	b1, err := plotter.NewBoxPlot(w, 1, normal)
// 	b1.FillColor = color.RGBA{127, 188, 165, 1}

// 	if err != nil {
// 		panic(err)
// 	}

// 	b2, err := plotter.NewBoxPlot(w, 2, expon)
// 	b2.FillColor = color.RGBA{127, 188, 165, 1}

// 	if err != nil {
// 		panic(err)
// 	}

// 	p.Add(b0, b1, b2)

// 	// Set the X axis of the plot to nominal with
// 	// the given names for x = 0, x = 1, x = 2.
// 	p.NominalX("Uniform\nDistribution", "Normal\nDistribution",
// 		"Exponential\nDistribution")

// 	if err := p.Save(3*vg.Inch, 4*vg.Inch,
// 		"GoNum-Plot-s01-07-A.png"); err != nil {

// 		panic(err)
// 	}
// }

package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
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

	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// Make boxes for our data and add them to the plot.
	err := plotutil.AddBoxPlots(p, vg.Points(20),
		"Uniform\nDistribution", uniform,
		"Normal\nDistribution", normal,
		"Exponential\nDistribution", expon)

	if err != nil {
		panic(err)
	}

	if err := p.Save(3*vg.Inch, 4*vg.Inch,
		"GoNum-Plot-s01-07-B.png"); err != nil {

		panic(err)
	}
}
