// package main

// import (
// 	"image/color"
// 	"math"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	p := plot.New()

// 	p.Title.Text = "Functions"
// 	p.X.Label.Text = "X"
// 	p.Y.Label.Text = "Y"

// 	// A quadratic function x^2
// 	quad := plotter.NewFunction(func(x float64) float64 { return x * x })
// 	quad.Color = color.RGBA{B: 255, A: 255}

// 	// An exponential function 2^x
// 	exp := plotter.NewFunction(func(x float64) float64 {
// 		return math.Pow(2, x) })
// 	exp.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
// 	exp.Width = vg.Points(2)
// 	exp.Color = color.RGBA{G: 255, A: 255}

// 	// The sine function, shifted and scaled
// 	// to be nicely visible on the plot.
// 	sin := plotter.NewFunction(func(x float64) float64 {
// 		return 10*math.Sin(x) + 50 })
// 	sin.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
// 	sin.Width = vg.Points(4)
// 	sin.Color = color.RGBA{R: 255, A: 255}

// 	// Add the funcitons and their legend entries.
// 	p.Add(quad, exp, sin)
// 	p.Legend.Add("x^2", quad)
// 	p.Legend.Add("2^x", exp)
// 	p.Legend.Add("10*sin(x) + 50", sin)
// 	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

// 	// Set the axis ranges. Unlike other data sets,
// 	// functions don't set the axis ranges automatically
// 	// since functions don't necessarily have a finite
// 	// range of x and y values.
// 	p.X.Min = 0
// 	p.X.Max = 10
// 	p.Y.Min = 0
// 	p.Y.Max = 100

// 	// Save the plot to a PNG file.
// 	if err := p.Save(4*vg.Inch, 4*vg.Inch,
// 		"GoNum-Plot-s01-06-A.png"); err != nil {

// 		panic(err)
// 	}
// }

package main

import (
	"image/color"
	// "math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/gonum/stat/distuv"
)

func main() {
	// Draw some random values from the standard
	// normal distribution.
	rand.Seed(int64(0))

	v := make(plotter.Values, 10_000)

	for i := range v {
		v[i] = rand.NormFloat64()
	}

	// Make a plot and set its title.
	p := plot.New()

	p.Title.Text = "Histogram"

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)

	if err != nil {
		panic(err)
	}

	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch,
		"GoNum-Plot-s01-06-A.png"); err != nil {

		panic(err)
	}
}
