package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/ajstarks/gpdf"
)

func main() {
	var output = "f.pdf"
	if len(os.Args) > 1 {
		output = os.Args[1]
	}
	ps := gpdf.Letter
	canvas, err := gpdf.SetupCanvas(ps)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// API
	//////////////////////////////////////////////////////////////////////////////
	labels := []string{

		"Arc(x,y, r,a1,a2, size,color)",
		"Circle(x,y, r, color)",
		"Ellipse(x,y, w,h, color)",
		"Image(x,y, w,h, name)",
		"Line(x1,y1, x2,y2, size, color)",
		"Polygon(x,y, color)",
		"QuadCurve(bx,by, cx,cy, ex,ey, size, color)",
		"Rect(x,y, w,h, color)",
		"Text(x,y, size, color)",
		"TextCode(name, x,y,w,h, size,border, bcolor,tcolor)",
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
	incr := 5.0
	ts := 2.5
	dotsize := 0.4
	shapecolor := "rgb(200,200,200)" //"lightgray"
	dotcolor := "white"
	tcolor := "gray"
	canvas.Background("black")
	canvas.Text(c1, 95, 3, "Generate PDF (gpdf) API", "white")
	canvas.Font = gpdf.Mono
	for i := range labels {
		canvas.Text(c1, y, ts, labels[i], tcolor)
		y -= yspace
	}
	y = 85.0
	// arc
	canvas.Circle(c2, y, dotsize, dotcolor)
	canvas.Arc(c2, y, 5, 0, 90, lw, shapecolor)
	y -= yspace
	// circle
	canvas.Circle(c2, y, 5, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// ellipse
	canvas.Ellipse(c2, y, 5, 2, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// image
	canvas.Image(c2, y, 64, 48, "follow.jpg")
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
	canvas.Polygon(xp, yp, shapecolor)
	y -= yspace
	// quadcurve
	bx := c3
	by := y
	cx := c2 + 5
	cy := y + yspace2
	ex := c4
	ey := y
	canvas.QuadCurve(bx, by, cx, cy, ex, ey, lw, shapecolor)
	canvas.Circle(bx, by, dotsize, dotcolor)
	canvas.Circle(cx, cy, dotsize, dotcolor)
	canvas.Circle(ex, ey, dotsize, dotcolor)
	y -= yspace
	// rect
	canvas.Rect(c2, y, 4, 5, shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// text
	canvas.Text(c2, y, 3, "hello", shapecolor)
	canvas.Circle(c2, y, dotsize, dotcolor)
	y -= yspace
	// textcode
	canvas.TextCode("hello.txt", c3, y+4, 12, yspace*0.6, 1.2, 0.3, shapecolor, "gray")
	canvas.Circle(c3, y+4, dotsize, dotcolor)
	y -= yspace
	// grid
	gy1 := y - 2.5
	gy2 := y + 2.5
	canvas.Grid(c3, c4, gy1, gy2, 0.1, 2.5, shapecolor)
	canvas.Circle(85, gy1, dotsize, dotcolor)
	canvas.Circle(95, gy2, dotsize, dotcolor)
	//////////////////////////////////////////////////////////////////////////////////////////////

	// usage
	canvas.Font = gpdf.Sans
	canvas.NewPage(ps)
	canvas.Background("black")

	for range 500 {
		xr := rand.Float64() * 100
		yr := rand.Float64() * 100
		canvas.Circle(xr, yr, 0.25, "white")
	}

	canvas.Circle(50, 0, 50, "blue")
	canvas.Text(25, 20, 10, "hello, world", "white")
	codesize := 1.3
	border := codesize / 8
	canvas.TextCode("excode", 2, 98, 42, 45, codesize, border, "black", "white")

	//////////////////////////////////////////////////////////////////////////////////////////////
	// play
	canvas.NewPage(ps)
	canvas.Grid(0, 100, 0, 100, 0.1, incr, "gray")
	for v := 0.0; v <= 100; v += incr {
		canvas.Circle(100-v, v, 1, "blue")
		canvas.Circle(v, v, 1, "red")
		canvas.Circle(v, 50, 1, "green")
	}
	canvas.Circle(50, 50, 12, "red")
	canvas.Rect(20, 80, 5, 5, "red")
	canvas.Rect(80, 20, 5, 5, "blue")
	canvas.Ellipse(80, 80, 5, 2.5, "blue")
	canvas.Ellipse(20, 20, 5, 2.5, "red")
	canvas.Circle(20, 80, 0.5, "white")
	canvas.Circle(80, 20, 0.5, "white")
	canvas.Circle(50, 50, 2, "white")
	canvas.Circle(80, 80, 0.5, "white")
	canvas.Circle(20, 20, 0.5, "white")
	canvas.Polygon([]float64{40, 50, 60}, []float64{5, 15, 5}, "green")
	canvas.QuadCurve(5, 55, 5, 60, 20, 55, 0.5, "red")
	canvas.QuadCurve(80, 55, 95, 60, 95, 55, 0.5, "red")
	canvas.Image(50, 90, 133, 100, "follow.jpg")

	err = canvas.Creator.WriteToFile(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
