package main

import (
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	lnPlotData := make(plotter.XYs, 100)

	x := 0.1

	for i := 0; i < 100; i++ {
		lnPlotData[i].X = x
		lnPlotData[i].Y = math.Log(x)

		x += 0.1
	}

	plotObj := plot.New()

	plotObj.Title.Text = "ln(x)"

	plotObj.X.Label.Text = "x"
	plotObj.Y.Label.Text = "y"

	err := plotutil.AddLinePoints(plotObj,
		"Natural logarithm", lnPlotData)

	if err != nil {
		panic(err)
	}

	if err := plotObj.Save(8*vg.Inch, 8*vg.Inch,
		"lnPlotPicture.png"); err != nil {
		panic(err)
	}
}
