package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/gpdf"
)

type Point struct{ X, Y float64 }

func api(canvas *gpdf.Canvas) {
	var colx float64
	var lw float64 = 0.2
	var labelsize float64 = 2.0
	titlesize := labelsize * 2
	subsize := labelsize * 0.7
	op := 30.0
	tcolor := "#aa0000"
	fcolor := "#0000aa"
	bgcolor := "white"
	labelcolor := "rgb(50,50,50)"
	subtitle := "A canvas API for PDF using high-level objects and a percentage-based coordinate system (https://github.com/ajstarks/gpdf)"

	// Title
	canvas.Background(bgcolor)
	colx = 20
	canvas.CText(colx, 92, titlesize, "gpdf API", labelcolor)
	canvas.TextWrap(colx+15, 95, 40, titlesize*0.3, 1.2, subtitle, labelcolor)

	// Lines
	canvas.CText(colx, 80, labelsize, "Line", labelcolor)
	canvas.Line(10, 70, colx+5, 65, lw, tcolor, op)
	canvas.Coord(10, 70, subsize, "P0", labelcolor)
	canvas.Coord(colx+5, 65, subsize, "P1", labelcolor)

	canvas.Line(colx, 70, 35, 75, lw, fcolor, op)
	canvas.Coord(colx, 70, subsize, "P0", labelcolor)
	canvas.Coord(35, 75, subsize, "P1", labelcolor)

	// Circle
	cx1 := colx - 10
	cx2 := colx + 10
	canvas.CText(cx1, 55, labelsize, "Circle", labelcolor)
	canvas.Circle(cx1, 45, 5, fcolor, op)
	canvas.Coord(cx1, 45, subsize, "center", labelcolor)

	// Arc
	canvas.CText(cx2, 55, labelsize, "Arc", labelcolor)
	canvas.Arc(cx2, 45, 5, 5, 0, 215, 5, tcolor, op)
	canvas.Coord(cx2, 45, subsize, "center", labelcolor)

	// Ellipse
	canvas.CText(colx, 30, labelsize, "Ellipse", labelcolor)
	canvas.Ellipse(colx, 15, 5, 10, tcolor, op)
	canvas.Ellipse(colx, 15, 10, 5, fcolor, op)
	canvas.Coord(colx, 15, subsize, "center", labelcolor)

	// Quadradic Bezier
	start := Point{X: 45, Y: 65}
	c1 := Point{X: 70, Y: 85}
	end := Point{X: 70, Y: 65}
	canvas.CText(60, 80, labelsize, "Quadratic Bezier Curve", labelcolor)
	canvas.Curve(start.X, start.Y, c1.X, c1.Y, end.X, end.Y, lw, tcolor)
	canvas.Coord(start.X, start.Y, subsize, "start", labelcolor)
	canvas.Coord(c1.X, c1.Y, subsize, "control", labelcolor)
	canvas.Coord(end.X, end.Y, subsize, "end", labelcolor)

	colx += 40
	// Cubic Bezier
	start = Point{X: 45, Y: 40}
	c1 = Point{X: 45, Y: 55}
	c2 := Point{X: colx, Y: 50}
	end = Point{X: 70, Y: 40}
	canvas.CText(colx, 55, labelsize, "Cubic Bezier Curve", labelcolor)
	canvas.CubicCurve(start.X, start.Y, c1.X, c1.Y, c2.X, c2.Y, end.X, end.Y, lw, fcolor)
	canvas.Coord(start.X, start.Y, subsize, "start", labelcolor)
	canvas.Coord(end.X, end.Y, subsize, "end", labelcolor)
	canvas.Coord(c1.X, c1.Y, subsize, "control 1", labelcolor)
	canvas.Coord(c2.X, c2.Y, subsize, "control 2", labelcolor)

	// Polygon
	canvas.CText(colx, 30, labelsize, "Polygon", labelcolor)
	xp := []float64{45, 60, 70, 70, 60, 45}
	yp := []float64{25, 20, 25, 5, 10, 5}
	for i := range xp {
		canvas.Coord(xp[i], yp[i], subsize, fmt.Sprintf("P%d", i), labelcolor)
	}
	canvas.Polygon(xp, yp, fcolor, op)

	colx += 30
	// Rectangles
	canvas.CText(colx, 80, labelsize, "Rectangle", labelcolor)
	canvas.Rect(colx, 70, 5, 15, fcolor, op)
	canvas.Coord(colx, 70, subsize, "center", labelcolor)

	// Square
	canvas.CText(colx, 55, labelsize, "Square", labelcolor)
	canvas.Square(colx, 45, 10, tcolor, op)
	canvas.Coord(colx, 45, subsize, "center", labelcolor)

	// Image
	canvas.CText(colx, 30, labelsize, "Image", labelcolor)
	canvas.ImageName(colx, 15, 75, 75, "earth.jpg")
	canvas.Coord(colx, 15, subsize, "", "black")

}

func main() {
	cw, ch := 800.0, 500.0
	canvas, err := gpdf.SetupCanvas(cw, ch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	font, err := canvas.LoadFontFile("sans.ttf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	canvas.SetFont(font)
	api(canvas)
	canvas.Creator.WriteToFile("api.pdf")

}
