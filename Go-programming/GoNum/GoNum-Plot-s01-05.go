// package main

// import (
// 	"math/rand"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	// Get some data.
// 	n, m := 5, 10
// 	pts := make([]plotter.XYer, n)

// 	for i := range pts {
// 		xys := make(plotter.XYs, m)
// 		pts[i] = xys
// 		center := float64(i)

// 		for j := range xys {
// 			xys[j].X = center + (rand.Float64() - 0.5)
// 			xys[j].Y = center + (rand.Float64() - 0.5)
// 		}
// 	}

// 	plt := plot.New()

// 	// Creat two lines connecting points and errors bars. For
// 	// the first, each point is the mean x and y value and the
// 	// error bars give the 95% confidence intervals. For the
// 	// second, each point is the median x and y value with the
// 	// error bars showing the minimum and maximum values.

// 	mean95, err := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pts...)

// 	if err != nil {
// 		panic(err)
// 	}

// 	medMinMax, err := plotutil.NewErrorPoints(plotutil.MedianAndMinMax,
// 		pts...)

// 	if err != nil {
// 		panic(err)
// 	}

// 	plotutil.AddLinePoints(plt,
// 		"mean and 95% confidence", mean95,
// 		"median and minimum and maximum", medMinMax)

// 	plotutil.AddErrorBars(plt, mean95, medMinMax)

// 	// Add the points that are summarized by the error points.
// 	plotutil.AddScatters(plt, pts[0], pts[1], pts[2], pts[3], pts[4])

// 	plt.Save(4*vg.Inch, 4*vg.Inch, "errpoints.png")
// }

package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	groupA := plotter.Values{20, 35, 30, 35, 27}
	groupB := plotter.Values{25, 32, 34, 20, 25}
	groupC := plotter.Values{12, 28, 15, 21, 8}

	p := plot.New()

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)

	if err != nil {
		panic(err)
	}

	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w



	barsB, err := plotter.NewBarChart(groupB, w)

	if err != nil {
		panic(err)
	}

	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(1)



	barsC, err := plotter.NewBarChart(groupC, w)

	if err != nil {
		panic(err)
	}

	barsC.LineStyle.Width = vg.Length(0)
	barsC.Color = plotutil.Color(2)
	barsC.Offset = w

	p.Add(barsA, barsB, barsC)
	p.Legend.Add("Group A", barsA)
	p.Legend.Add("Group B", barsB)
	p.Legend.Add("Gropu C", barsC)
	p.Legend.Top = true
	p.NominalX("One", "Two", "Three", "Four", "Five")

	if err := p.Save(5*vg.Inch, 3*vg.Inch,
		"GoNume-Plot-s01-05-A.png"); err != nil {
		panic(err)
	}

}
