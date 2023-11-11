// package main

// import (
// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
// )

// func main() {
// 	pointsForPlot := make(plotter.XYs, 10)

// 	pointsForPlot[0].X = 0.0
// 	pointsForPlot[0].Y = 1.0

// 	pointsForPlot[1].X = 1.0
// 	pointsForPlot[1].Y = 0.1

// 	pointsForPlot[2].X = 2.0
// 	pointsForPlot[2].Y = -0.1

// 	pointsForPlot[3].X = 3.0
// 	pointsForPlot[3].Y = 0.05

// 	pointsForPlot[4].X = 4.0
// 	pointsForPlot[4].Y = 0.1

// 	pointsForPlot[5].X = 5.0
// 	pointsForPlot[5].Y = -0.05

// 	pointsForPlot[6].X = 6.0
// 	pointsForPlot[6].Y = 0.5

// 	pointsForPlot[7].X = 7.0
// 	pointsForPlot[7].Y = 0.75

// 	pointsForPlot[8].X = 8.0
// 	pointsForPlot[8].Y = 1.2

// 	pointsForPlot[9].X = 9.0
// 	pointsForPlot[9].Y = 0.8



// 	plotVar := plot.New()

// 	plotVar.Title.Text = "Test robienia wykresow"

// 	plotVar.X.Label.Text = "Czas"
// 	plotVar.Y.Label.Text = "Masa [kg]"

// 	err := plotutil.AddLinePoints(plotVar,
// 		"Test robienia wykresow", pointsForPlot)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := plotVar.Save(7*vg.Inch, 7*vg.Inch,
// 		"TestPlot.png"); err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"image/color"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	// Get some random points.
	rand.Seed(int64(0))

	n := 15
	scatterData := randomPoints(n)
	lineData := randomPoints(n)
	linePointsData := randomPoints(n)

	// Create a new plot, set its title and axis labels.
	p := plot.New()

	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	// Draw a grid behind the data.
	p.Add(plotter.NewGrid())

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(scatterData)

	if err != nil {
		panic(err)
	}

	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	// Make a line plotter and set its style.
	l, err := plotter.NewLine(lineData)

	if err != nil {
		panic(err)
	}

	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Make a line plotter with points and set its style.
	lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)

	if err != nil {
		panic(err)
	}

	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.PyramidGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A:255}

	// Add the plotters to the plot, with a legend entry for each.
	p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	p.Legend.Add("line", l)
	p.Legend.Add("line points", lpLine, lpPoints)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch,
		"GoNum-Plot-s01-02-A.png"); err != nil {

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

		pts[i].Y = pts[i].X + 10*rand.Float64()
	}

	return pts
}
