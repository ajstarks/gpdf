package gpdf

import (
	"strconv"
)

// pct returns the percentage of its input
func pct(p float64, m float64) float64 {
	return ((p / 100.0) * m)
}

// dimen returns canvas dimensions from percentages
// (converting from x increasing left-right, y increasing top-bottom)
func dimen(xp, yp, w, h float64) (float64, float64) {
	return pct(xp, w), pct(yp, h)
}

// Line strokes a colored line from (x0, y0) to (x1, y1)
func (c *Canvas) Line(x0, y0, x1, y1, size float64, color string) {
	x0, y0 = dimen(x0, y0, c.Width, c.Height)
	x1, y1 = dimen(x1, y1, c.Width, c.Height)
	size = pct(size, c.Width)
	c.AbsLine(x0, y0, x1, y1, size, color)
}

// Circle makes a color filled circle centered at (x, y) with radius r
func (c *Canvas) Circle(x, y, r float64, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	r = pct(r, c.Width)
	c.AbsCircle(x, y, r, color)
}

// Ellipse draws an ellipse centered at (x, y) with dimensions (w,h)
func (c *Canvas) Ellipse(x, y, w, h float64, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	c.AbsEllipse(x, y, w, h, color)
}

// Rect makes a color filled rectangle centered at (x,y), with dimentions (w,h)
func (c *Canvas) Rect(x, y, w, h float64, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	c.AbsCenterRect(x, y, w, h, color)
}

// Text places text at (x,y) at the specified size and color
func (c *Canvas) Text(x, y, size float64, s string, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	c.AbsText(x, y, size, s, color)
}

// Polygon makes a color filled polygon using the specified coordinates in x and y
func (c *Canvas) Polygon(x, y []float64, color string) {
	for i := range x {
		x[i], y[i] = dimen(x[i], y[i], c.Width, c.Height)
	}
	c.AbsPolygon(x, y, color)
}

// QuadCurve draws a quadradic bezier curve
// begin coordinates (bx, by)
// control coordinates (cx, cy)
// end coordinates at (ex, ey)
func (c *Canvas) QuadCurve(bx, by, cx, cy, ex, ey, size float64, color string) {
	bx, by = dimen(bx, by, c.Width, c.Height)
	cx, cy = dimen(cx, cy, c.Width, c.Height)
	ex, ey = dimen(ex, ey, c.Width, c.Height)
	size = pct(size, c.Width)
	c.AbsCubicBezier(bx, by, cx, cy, cx, cy, ex, ey, size, color)
}

func (c *Canvas) Arc(x, y, w, h, a1, a2, size float64, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(w, c.Height)
	a1 = deg2rad(a1)
	a2 = deg2rad(a2)
	c.AbsArc(x, y, w, h, a1, a2, size, color)
}

// Images places an image centered at (x,y), with dimensions (w,h)
func (c *Canvas) Image(x, y, w, h float64, name string) {
	x, y = dimen(x, y, c.Width, c.Height)
	err := c.AbsCenterImage(x, y, w, h, name)
	if err != nil {
		c.Rect(x, y, w, h, "black")
	}
}

// PlainGrid makes a unlabeled coordinate grid starting at (xmin, ymin)
// ending at (xmax, ymax), line width is size, incr is the size of the grid
func (c *Canvas) Grid(xmin, xmax, ymin, ymax, size, incr float64, color string) {
	for v := xmin; v <= xmax; v += incr {
		c.Line(v, ymin, v, ymax, size, color)
	}
	for v := ymin; v <= ymax; v += incr {
		c.Line(xmin, v, xmax, v, size, color)
	}
}

// Grid makes a labeled coordinate grid starting at (xmin, ymin)
// ending at (xmax, ymax), line width is size, incr is the size of the grid
func (c *Canvas) LGrid(xmin, xmax, ymin, ymax, size, incr float64, color string) {
	for v := xmin; v <= xmax; v += incr {
		c.Line(v, ymin, v, ymax, size, color)
		c.Text(v, ymin, incr*.4, strconv.FormatFloat(v, 'g', -1, 64), color)
	}
	for v := ymin; v <= ymax; v += incr {
		c.Line(xmin, v, xmax, v, size, color)
		c.Text(xmin, v, incr*.4, strconv.FormatFloat(v, 'g', -1, 64), color)
	}
}
