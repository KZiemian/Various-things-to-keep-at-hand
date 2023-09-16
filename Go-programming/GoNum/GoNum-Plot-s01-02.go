package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	pointsForPlot := make(plotter.XYs, 10)

	pointsForPlot[0].X = 0.0
	pointsForPlot[0].Y = 1.0

	pointsForPlot[1].X = 1.0
	pointsForPlot[1].Y = 0.1

	pointsForPlot[2].X = 2.0
	pointsForPlot[2].Y = -0.1

	pointsForPlot[3].X = 3.0
	pointsForPlot[3].Y = 0.05

	pointsForPlot[4].X = 4.0
	pointsForPlot[4].Y = 0.1

	pointsForPlot[5].X = 5.0
	pointsForPlot[5].Y = -0.05

	pointsForPlot[6].X = 6.0
	pointsForPlot[6].Y = 0.5

	pointsForPlot[7].X = 7.0
	pointsForPlot[7].Y = 0.75

	pointsForPlot[8].X = 8.0
	pointsForPlot[8].Y = 1.2

	pointsForPlot[9].X = 9.0
	pointsForPlot[9].Y = 0.8



	plotVar := plot.New()

	plotVar.Title.Text = "Test robienia wykresow"

	plotVar.X.Label.Text = "Czas"
	plotVar.Y.Label.Text = "Masa [kg]"

	err := plotutil.AddLinePoints(plotVar,
		"Test robienia wykresow", pointsForPlot)

	if err != nil {
		panic(err)
	}

	if err := plotVar.Save(7*vg.Inch, 7*vg.Inch,
		"TestPlot.png"); err != nil {
		panic(err)
	}
}
