// package main

// import (
// 	"math"
// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	lnPlotData := make(plotter.XYs, 100)

// 	x := 0.1

// 	for i := 0; i < 100; i++ {
// 		lnPlotData[i].X = x
// 		lnPlotData[i].Y = math.Log(x)

// 		x += 0.1
// 	}

// 	plotObj := plot.New()

// 	plotObj.Title.Text = "ln(x)"

// 	plotObj.X.Label.Text = "x"
// 	plotObj.Y.Label.Text = "y"

// 	err := plotutil.AddLinePoints(plotObj,
// 		"Natural logarithm", lnPlotData)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := plotObj.Save(8*vg.Inch, 8*vg.Inch,
// 		"lnPlotPicture.png"); err != nil {
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
	rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = "Relabeling ticks marks example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	// Use a custom tick marker interface implementation with the Ticks
	// function, that computers the default tick marks and re-labels
	// the major ticks with commas.
	p.Y.Tick.Marker = commaTicks{}

	err := plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))

	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch,
		"GoNum-s01-04-A.png"); err != nil {

		panic(err)
	}
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)

	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i - 1].X + rand.Float64()
		}

		pts[i].Y = (pts[i].X + 10*rand.Float64()) * 1000
	}

	return pts
}

type commaTicks struct{}

// Ticks copmutes the default tick marks, but inserts commas
// into the labels for the major tick marks.
func (commaTicks) Ticks(min, max float64) []plot.Tick {
	tks := plot.DefaultTicks{}.Ticks(min, max)

	for i, t := range tks {
		if t.Label == "" { // Skip minor ticks, they are fine.
			continue
		}

		tks[i].Label = addCommas(t.Label)
	}

	return tks
}

// addCommas adds commas after every 3 characters from right to left.
// NOTE: This funciton is a quick hack, it doesn't work with decimal
// points, and may have a bunch of other problems.
func addCommas(s string) string {
	rev := ""
	n := 0

	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
		n++

		if (n % 3) == 0 {
			rev += ","
		}
	}

	s = ""

	for i := len(rev) - 1; i >= 0; i-- {
		s += string(rev[i])
	}

	return s
}
