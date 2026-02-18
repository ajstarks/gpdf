package gpdf

import (
	"bufio"
	"math"
	"os"
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

// Background set the page background color
func (c *Canvas) Background(color string) {
	c.Rect(50, 50, 100, 100, color)
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

// CornerRect a color filled rectangle lower left at (x,y), with dimentions (w,h)
func (c *Canvas) CornerRect(x, y, w, h float64, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	c.AbsCornerRect(x, y-h, w, h, color)
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

// BText places text begin aligned at (x,y) at the specified size and color
func (c *Canvas) BText(x, y, size float64, s string, color string) {
	c.Text(x, y, size, s, color)
}

// CText places text centered at (x,y) at the specified size and color
func (c *Canvas) CText(x, y, size float64, s string, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	c.AbsCText(x, y, size, s, color)
}

// EText places text end-aligned at (x,y) at the specified size and color
func (c *Canvas) EText(x, y, size float64, s string, color string) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	c.AbsEText(x, y, size, s, color)
}

// TextCode shows a text file, upper left at (x,y), dimensions (w, h) at size
// the border of the block background and textcolors are also specified.
func (c *Canvas) TextCode(name string, x, y, w, h, size, border float64, bgcolor, txcolor string) {
	cf := c.CustomFont
	if c.CustomFont != nil {
		cf = c.CustomFont
	}
	c.CustomFont = nil
	pf := c.StdFont
	c.StdFont = Mono
	c.CornerRect(x-border, y+border, w+(border*2), h+(border*2), txcolor)
	c.CornerRect(x, y, w, h, bgcolor)
	r, err := os.Open(name)
	if err != nil {
		c.CText(x+(size/2), y-(size/2), size, "File not found", "red")
		return
	}
	scanner := bufio.NewScanner(r)
	y -= size
	for scanner.Scan() {
		c.Text(x+size, y, size, scanner.Text(), txcolor)
		y -= size * 1.2
	}
	c.StdFont = pf
	c.CustomFont = cf
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

// Arc makes a stroked arc, using percentage-based measures
// center is (x, y), the arc begins at angle a1, and ends at a2, with radius r.
// The arc is stroked with the specified stroke size and color
func (c *Canvas) Arc(x, y, r, a1, a2, size float64, fillcolor string) {
	// Define minimum and maximum step sizes
	const minstep = 0.001
	const maxstep = 0.1
	const twoPi = math.Pi * 2

	// convert angles from degrees to radians
	a1 = a1 * (math.Pi / 180)
	a2 = a2 * (math.Pi / 180)

	// Ensure the angles are in the range [0, 2π)
	a1 = math.Mod(a1, twoPi)
	a2 = math.Mod(a2, twoPi)
	// Calculate step size based on the radius (Smaller steps for larger radius)
	step := 1.0 / (3.0 * r * twoPi)

	// Clamp step to be within the defined range for performance reasons
	if step < minstep {
		step = minstep
	}
	if step > maxstep {
		step = maxstep
	}
	// Ensure we handle crossing the 0/2π boundary correctly
	if a2 < a1 {
		a2 += twoPi
	}
	// Initialize the starting point
	x1, y1 := c.Polar(x, y, r, a1)

	for t := a1; t < a2; t += step {
		x2, y2 := c.Polar(x, y, r, t)
		c.Line(x1, y1, x2, y2, size, fillcolor)
		x1 = x2
		y1 = y2
	}
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
