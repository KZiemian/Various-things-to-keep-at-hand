package main

import (
	"image/color"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(int64(0))

	n := 10
	bubbleData := randomTriples(n)

	p := plot.New()

	p.Title.Text = "Bubbles"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, err := plotter.NewBubbles(bubbleData, vg.Points(1), vg.Points(20))

	if err != nil {
		panic(err)
	}

	bs.Color = color.RGBA{R: 196, B: 128, A: 255}

	p.Add(bs)

	if err := p.Save(4*vg.Inch, 4*vg.Inch,
		"GoNum-Plot-s01-09.png"); err != nil {

		panic(err)
	}
}

// randomTriples returns some random x, y, z triples
// with some interesting kind of trend.
func randomTriples(n int) plotter.XYZs {
	data := make(plotter.XYZs, n)

	for i := range data {
		if i == 0 {
			data[i].X = rand.Float64()
		} else {
			data[i].X = data[i - 1].X + 2*rand.Float64()
		}

		data[i].Y = data[i].X + 10*rand.Float64()
		data[i].Z = data[i].X
	}

	return data
}
