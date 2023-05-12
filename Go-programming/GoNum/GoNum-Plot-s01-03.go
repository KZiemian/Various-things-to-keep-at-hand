package main

import (
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	lnPlotData := make(plotter.XYs, 45)

	lnPlotData[0].X = 0.1
	lnPlotData[0].Y = math.Log(0.1)

	lnPlotData[1].X = 0.2
	lnPlotData[1].Y = math.Log(0.2)

	lnPlotData[2].X = 0.3
	lnPlotData[2].Y = math.Log(0.3)

	lnPlotData[3].X = 0.4
	lnPlotData[3].Y = math.Log(0.4)

	lnPlotData[4].X = 0.5
	lnPlotData[4].Y = math.Log(0.5)

	lnPlotData[5].X = 0.6
	lnPlotData[5].Y = math.Log(0.6)

	lnPlotData[6].X = 0.7
	lnPlotData[6].Y = math.Log(0.7)

	lnPlotData[7].X = 0.8
	lnPlotData[7].Y = math.Log(0.8)

	lnPlotData[8].X = 0.9
	lnPlotData[8].Y = math.Log(0.9)

	lnPlotData[9].X = 1.0
	lnPlotData[9].Y = math.Log(1.0)

	lnPlotData[10].X = 1.1
	lnPlotData[10].Y = math.Log(1.1)

	lnPlotData[11].X = 1.2
	lnPlotData[11].Y = math.Log(1.2)

	lnPlotData[12].X = 1.3
	lnPlotData[12].Y = math.Log(1.3)

	lnPlotData[13].X = 1.4
	lnPlotData[13].Y = math.Log(1.4)

	lnPlotData[14].X = 1.5
	lnPlotData[14].Y = math.Log(1.5)

	lnPlotData[15].X = 1.6
	lnPlotData[15].Y = math.Log(1.6)

	lnPlotData[16].X = 1.7
	lnPlotData[16].Y = math.Log(1.7)

	lnPlotData[17].X = 1.8
	lnPlotData[17].Y = math.Log(1.8)

	lnPlotData[18].X = 1.9
	lnPlotData[18].Y = math.Log(1.9)

	lnPlotData[19].X = 2.0
	lnPlotData[19].Y = math.Log(2.0)

	lnPlotData[20].X = 2.1
	lnPlotData[20].Y = math.Log(2.1)

	lnPlotData[21].X = 2.2
	lnPlotData[21].Y = math.Log(2.2)

	lnPlotData[22].X = 2.3
	lnPlotData[22].Y = math.Log(2.3)

	lnPlotData[23].X = 2.4
	lnPlotData[23].Y = math.Log(2.4)

	lnPlotData[24].X = 2.5
	lnPlotData[24].Y = math.Log(2.5)

	lnPlotData[25].X = 2.6
	lnPlotData[25].Y = math.Log(2.6)

	lnPlotData[26].X = 2.7
	lnPlotData[26].Y = math.Log(2.7)

	lnPlotData[27].X = 2.8
	lnPlotData[27].Y = math.Log(2.8)

	lnPlotData[28].X = 2.9
	lnPlotData[28].Y = math.Log(2.9)

	lnPlotData[29].X = 3.0
	lnPlotData[29].Y = math.Log(3.0)

	lnPlotData[30].X = 3.1
	lnPlotData[30].Y = math.Log(3.1)

	lnPlotData[31].X = 3.2
	lnPlotData[31].Y = math.Log(3.2)

	lnPlotData[32].X = 3.3
	lnPlotData[32].Y = math.Log(3.3)

	lnPlotData[33].X = 3.4
	lnPlotData[33].Y = math.Log(3.4)

	lnPlotData[34].X = 3.5
	lnPlotData[34].Y = math.Log(3.5)

	lnPlotData[35].X = 3.6
	lnPlotData[35].Y = math.Log(3.6)

	lnPlotData[36].X = 3.7
	lnPlotData[36].Y = math.Log(3.7)

	lnPlotData[37].X = 3.8
	lnPlotData[37].Y = math.Log(3.8)

	lnPlotData[38].X = 3.9
	lnPlotData[38].Y = math.Log(3.9)

	lnPlotData[39].X = 4.0
	lnPlotData[39].Y = math.Log(4.0)

	lnPlotData[40].X = 4.1
	lnPlotData[40].Y = math.Log(4.1)

	lnPlotData[41].X = 4.2
	lnPlotData[41].Y = math.Log(4.2)

	lnPlotData[42].X = 4.3
	lnPlotData[42].Y = math.Log(4.3)

	lnPlotData[43].X = 4.4
	lnPlotData[43].Y = math.Log(4.4)

	lnPlotData[44].X = 10.0
	lnPlotData[44].Y = math.Log(10.0)



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
		"lnPlot.png"); err != nil {
		panic(err)
	}
}
