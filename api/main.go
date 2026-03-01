package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/gpdf"
)

type Point struct{ X, Y float64 }

// apitable makes a table of API calls
func apitable(canvas *gpdf.Canvas) {
	// API
	labels := []string{
		"Arc(x,y, w,h,a1,a2, size,color)",
		"Circle(x,y, r, color)",
		"Ellipse(x,y, w,h, color)",
		"ImageName(x,y, w,h, name)",
		"Line(x1,y1, x2,y2, size, color)",
		"Polygon(x,y, color), Polyline(x,y size, color)",
		"QuadCurve(bx,by, cx,cy, ex,ey, size, color)",
		"Square(x,y, w, color), Rect(x, y, w,h, color), GradRect(x,y,w,h,c1,c2,pct)",
		"{B,C,E}Text(x,y, size, s, color), RText(x,y, size,angle, s, color)",
		"TextWrap(x,y,w,size,linespacing,s,color)",
		"Grid(x1,x2, y1,y2, size, incr, color)",
		"Background(color)",
	}
	yspace := 7.5
	yspace2 := yspace / 2
	y := 85.0
	lw := 0.3
	c1 := 5.0
	c2 := c1 + 85
	c3 := c2 - 5
	c4 := c2 + 5
	ts := 1.8
	dotsize := 0.4
	shapecolor := "rgb(180,180,180)" //"lightgray"
	dotcolor := "black"
	tcolor := "gray"
	bgcolor := "white"
	fgcolor := dotcolor
	canvas.Background(bgcolor)
	canvas.Text(c1, 92, 4, "Generate PDF (gpdf) API", fgcolor)
	for i := range labels {
		canvas.Text(c1, y, ts, labels[i], tcolor)
		y -= yspace
	}
	y = 85.0
	// arc
	canvas.Circle(c2, y, dotsize, dotcolor)
	canvas.Arc(c2, y, 5, 5, 0, 90, lw, shapecolor)

	y -= yspace
	// circle
	canvas.Circle(c2, y, 3, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)

	y -= yspace
	// ellipse
	canvas.Ellipse(c2, y, 5, 2, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)

	y -= yspace
	// image
	canvas.ImageName(c2, y, 64, 48, "follow.jpg")
	canvas.Circle(c2, y, dotsize, dotcolor)

	y -= yspace
	// line
	canvas.Line(c3, y, c4, y, lw, shapecolor)
	canvas.Circle(c3, y, dotsize, dotcolor)
	canvas.Circle(c4, y, dotsize, dotcolor)

	y -= yspace
	// polygon
	xp := []float64{c3, c2, c4}
	yp := []float64{y, y + yspace2, y}
	canvas.Circle(xp[0], yp[0], dotsize, dotcolor)
	canvas.Circle(xp[1], yp[1], dotsize, dotcolor)
	canvas.Circle(xp[2], yp[2], dotsize, dotcolor)
	canvas.Polyline(xp, yp, 0.2, shapecolor)
	xp[0] -= 15.0
	xp[1] -= 15.0
	xp[2] -= 15.0
	canvas.Circle(xp[0], yp[0], dotsize, dotcolor)
	canvas.Circle(xp[1], yp[1], dotsize, dotcolor)
	canvas.Circle(xp[2], yp[2], dotsize, dotcolor)
	canvas.Polygon(xp, yp, shapecolor)

	y -= yspace
	// quadcurve
	bx := c3
	by := y
	cx := c2 + 5
	cy := y + yspace2
	ex := c4
	ey := y
	canvas.Curve(bx, by, cx, cy, ex, ey, lw, shapecolor)
	canvas.Circle(bx, by, dotsize, dotcolor)
	canvas.Circle(cx, cy, dotsize, dotcolor)
	canvas.Circle(ex, ey, dotsize, dotcolor)

	y -= yspace
	// rect
	canvas.Square(c2-10, y, 5, shapecolor)
	canvas.Circle(c2-10, y, dotsize, dotcolor)
	canvas.Rect(c2, y, 7, 5, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)

	y -= yspace
	// text
	tc0 := c2 - 18
	tc1 := c2 - 9
	tc2 := c2
	tc3 := c2 + 3
	canvas.BText(tc0, y, 2, "hello", shapecolor)
	canvas.Circle(tc0, y, dotsize, dotcolor)
	canvas.CText(tc1, y, 2, "hello", shapecolor)
	canvas.Circle(tc1, y, dotsize, dotcolor)
	canvas.EText(tc2, y, 2, "hello", shapecolor)
	canvas.Circle(tc2, y, dotsize, dotcolor)
	canvas.RText(tc3, y, 2, 45, "hello", shapecolor)
	canvas.Circle(tc3, y, dotsize, dotcolor)

	y -= yspace
	// Textwrap
	canvas.TextWrap(tc0, y, 25, 1.5, 1.2, "Now is the time for all good men to come to the aid of the party", shapecolor)

	y -= yspace
	// grid
	gy1 := y - 2.5
	gy2 := y + 2.5
	canvas.Grid(c3, c4, gy1, gy2, 0.1, 2.5, shapecolor)
	canvas.Circle(85, gy1, dotsize, dotcolor)
	canvas.Circle(95, gy2, dotsize, dotcolor)
}

// api makes an API placemat
func api(canvas *gpdf.Canvas) {
	colx := 20.0
	lw := 0.2
	labelsize := 2.0
	titlesize := labelsize * 2
	subsize := labelsize * 0.7
	op := 30.0
	top := 90.0
	tcolor := "rgb(128,0,0)"
	fcolor := "rgb(0,0,128)"
	bgcolor := "rgb(255,255,255)"
	labelcolor := "rgb(50,50,50)"
	subtitle := "A canvas API for PDF using high-level objects and a percentage-based coordinate system (https://github.com/ajstarks/gpdf)"

	// Title
	canvas.Background(bgcolor)

	canvas.CText(colx, top, titlesize, "gpdf API", labelcolor)
	canvas.TextWrap(colx+15, top+2, 30, titlesize*0.3, 1.2, subtitle, labelcolor)

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
	canvas.Curve(start.X, start.Y, c1.X, c1.Y, end.X, end.Y, lw, tcolor, op)
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
	canvas.CubicCurve(start.X, start.Y, c1.X, c1.Y, c2.X, c2.Y, end.X, end.Y, lw, fcolor, op)
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
	canvas.Coord(colx, 15, subsize, "", "white")

	// Text
	tx1 := colx - 5
	tx2 := colx
	tx3 := colx + 5
	tx4 := tx3 + 3
	ss := subsize * .6
	canvas.CText(colx, top+5, labelsize, "Text", labelcolor)
	canvas.Text(tx1, top, subsize, "hello", labelcolor)
	canvas.CText(tx2, top, subsize, "hello", labelcolor)
	canvas.EText(tx3, top, subsize, "hello", labelcolor)
	canvas.RText(tx4, top, subsize, 90, "hello", labelcolor)

	canvas.Square(tx1, top, ss, labelcolor, 50)
	canvas.Square(tx2, top, ss, labelcolor, 50)
	canvas.Square(tx3, top, ss, labelcolor, 50)
	canvas.Square(tx4, top, ss, labelcolor, 50)
	canvas.CText(tx1, top-2, ss, "begin", labelcolor)
	canvas.CText(tx2, top-2, ss, "center", labelcolor)
	canvas.CText(tx3, top-2, ss, "end", labelcolor)
	canvas.CText(tx4, top-2, ss, "rotate", labelcolor)
}

func main() {
	cw := 792.0 // 800.0
	ch := 612.0 // 500.0
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
	canvas.Creator.SetTitle("gpdf API")
	canvas.Creator.SetAuthor("Anthony Starks")
	canvas.SetFont(font)
	api(canvas)
	canvas.NewPage(cw, ch)
	apitable(canvas)
	canvas.Creator.WriteToFile("api.pdf")

}
