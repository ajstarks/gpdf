package gpdf

import (
	"math"
	"strconv"
	"strings"
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
func (c *Canvas) Background(fillcolor string) {
	color := ColorLookup(fillcolor)
	x, y := dimen(50, 50, c.Width, c.Height)
	w, h := pct(100, c.Width), pct(100, c.Height)
	c.AbsCenterRect(x, y, w, h, color)
}

// Line strokes a colored line from (x0, y0) to (x1, y1)
func (c *Canvas) Line(x0, y0, x1, y1, size float64, strokecolor string, opacity ...float64) {
	x0, y0 = dimen(x0, y0, c.Width, c.Height)
	x1, y1 = dimen(x1, y1, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(strokecolor)
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsLine(x0, y0, x1, y1, size, color)
}

// Circle makes a color filled circle centered at (x, y) with radius r
func (c *Canvas) Circle(x, y, r float64, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	r = pct(r, c.Width)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsCircle(x, y, r, color)
}

// Ellipse draws an ellipse centered at (x, y) with dimensions (w,h)
func (c *Canvas) Ellipse(x, y, w, h float64, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsEllipse(x, y, w, h, color)
}

// CornerRect a color filled rectangle lower left at (x,y), with dimentions (w,h)
func (c *Canvas) CornerRect(x, y, w, h float64, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsCornerRect(x, y-h, w, h, color)
}

// Rect makes a color filled rectangle centered at (x,y), with dimentions (w,h)
func (c *Canvas) Rect(x, y, w, h float64, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsCenterRect(x, y, w, h, color)
}

func (c *Canvas) GradRect(x, y, w, h float64, color1 string, color2 string, percent float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Width)
	h = pct(h, c.Height)
	c1 := ColorLookup(color1)
	c2 := ColorLookup(color2)
	c.AbsGradRect(x, y, w, h, c1, c2, percent)
}

// Square makes a square centered at (x,y), at size w
func (c *Canvas) Square(x, y, w float64, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	w = pct(w, c.Height)
	h := pct(100, w)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsCenterRect(x, y, w, h, color)
}

// Text places text at (x,y) at the specified size and color
func (c *Canvas) Text(x, y, size float64, s string, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsText(x, y, size, s, color)
}

// BText places text begin aligned at (x,y) at the specified size and color
func (c *Canvas) BText(x, y, size float64, s string, fillcolor string, opacity ...float64) {
	c.Text(x, y, size, s, fillcolor)
}

// CText places text centered at (x,y) at the specified size and color
func (c *Canvas) CText(x, y, size float64, s string, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsCText(x, y, size, s, color)
}

// EText places text end-aligned at (x,y) at the specified size and color
func (c *Canvas) EText(x, y, size float64, s string, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsEText(x, y, size, s, color)
}

// RText places text end-aligned at (x,y) at the specified size and color
func (c *Canvas) RText(x, y, size, angle float64, s string, fillcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsRText(x, y, size, angle, s, color)
}

func (c *Canvas) TextWrap(x, y, w, size, linespacing float64, s string, textcolor string, opacity ...float64) {
	x, y = dimen(x, y, c.Width, c.Height)
	size = pct(size, c.Width)
	w = pct(w, c.Width)
	color := ColorLookup(textcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	const factor = 0.3
	wordspacing := c.CustomFont.MeasureString("M", size) * factor
	xp := x
	yp := y
	edge := x + w
	words := strings.FieldsFunc(s, whitespace)
	//  never go over the edge
	for _, s := range words {
		tw := c.CustomFont.MeasureString(s, size)
		if xp+tw > edge {
			xp = x
			yp -= (size * linespacing)
		}
		c.AbsText(xp, yp, size, s, color)
		xp += tw + wordspacing
	}
}

// Polygon makes a color filled polygon using the specified coordinates in x and y
func (c *Canvas) Polygon(x, y []float64, fillcolor string, opacity ...float64) {
	for i := range x {
		x[i], y[i] = dimen(x[i], y[i], c.Width, c.Height)
	}
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsPolygon(x, y, color)
}

// Polyline makes connected lines using coordinates in x and y
func (c *Canvas) Polyline(x, y []float64, size float64, strokecolor string, opacity ...float64) {
	lx := len(x)
	if lx < 3 || lx != len(y) {
		return
	}
	for i := 0; i < lx-1; i++ {
		c.Line(x[i], y[i], x[i+1], y[i+1], size, strokecolor)
	}
}

// CubicCurve draws a quadradic bezier curve
// begin coordinates (bx, by)
// control coordinates (cx, cy)
// end coordinates at (ex, ey)
func (c *Canvas) CubicCurve(bx, by, cx1, cy1, cx2, cy2, ex, ey, size float64, strokecolor string, opacity ...float64) {
	bx, by = dimen(bx, by, c.Width, c.Height)
	cx1, cy1 = dimen(cx1, cy1, c.Width, c.Height)
	cx2, cy2 = dimen(cx2, cy2, c.Width, c.Height)
	ex, ey = dimen(ex, ey, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(strokecolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsCubicBezier(bx, by, cx1, cy1, cx2, cy2, ex, ey, size, color)
}

// QuadCurve draws a quadradic bezier curve
// begin coordinates (bx, by)
// control coordinates (cx, cy)
// end coordinates at (ex, ey)
func (c *Canvas) QuadCurve(bx, by, cx, cy, ex, ey, size float64, strokecolor string, opacity ...float64) {
	bx, by = dimen(bx, by, c.Width, c.Height)
	cx, cy = dimen(cx, cy, c.Width, c.Height)
	ex, ey = dimen(ex, ey, c.Width, c.Height)
	size = pct(size, c.Width)
	color := ColorLookup(strokecolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsQuadBezier(bx, by, cx, cy, ex, ey, size, color)
}

// Curve makes a quadradic bezier curve
func (c *Canvas) Curve(bx, by, cx, cy, ex, ey, size float64, strokecolor string, opacity ...float64) {
	c.QuadCurve(bx, by, cx, cy, ex, ey, size, strokecolor)
}

// Arc makes a stroked arc, using percentage-based measures
// center is (x, y), the arc begins at angle a1, and ends at a2, with radius r.
// The arc is stroked with the specified stroke size and color
func (c *Canvas) Arc(x, y, w, h, a1, a2, size float64, fillcolor string, opacity ...float64) {
	cw := c.Width
	ch := c.Height
	x, y = dimen(x, y, cw, ch)
	if w == h { // circular arc
		w = pct(w, ch)
		h = pct(100, w)
	} else { // ellipitcal arc
		h = pct(w, cw)
		w = pct(h, ch)
	}
	size = pct(size, cw)
	color := ColorLookup(fillcolor)
	color.A = 1
	if len(opacity) > 0 {
		color.A = opacity[0] / 100
	}
	c.AbsArc(x, y, w, h, a1, a2, size, color)
}

func (c *Canvas) OldArc(x, y, r, a1, a2, size float64, strokecolor string, opacity ...float64) {
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
		c.Line(x1, y1, x2, y2, size, strokecolor)
		x1 = x2
		y1 = y2
	}
}

// ImageName places a named image centered at (x,y), with dimensions (w,h)
func (c *Canvas) ImageName(x, y, w, h float64, name string) {
	x, y = dimen(x, y, c.Width, c.Height)
	err := c.AbsCenterImageName(x, y, w, h, name)
	if err != nil {
		c.Rect(x, y, w, h, "black")
	}
}

// Image places an image centered at (x,y), with dimensions (w,h)
func (c *Canvas) Image(x, y, w, h float64, img Pimage) {
	x, y = dimen(x, y, c.Width, c.Height)
	err := c.AbsCenterImage(x, y, w, h, img.Image)
	if err != nil {
		c.Rect(x, y, w, h, "black")
	}
}

// Grid makes a unlabeled coordinate grid starting at (xmin, ymin)
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
func (c *Canvas) LGrid(xmin, xmax, ymin, ymax, size, incr float64, color string, opacity ...float64) {
	var op float64
	if len(opacity) > 0 {
		op = opacity[0]
	}
	for v := xmin; v <= xmax; v += incr {
		c.Line(v, ymin, v, ymax, size, color, op)
		c.CText(v, ymin, incr*.3, strconv.FormatFloat(v, 'g', -1, 64), color, op)
	}
	for v := ymin; v <= ymax; v += incr {
		c.Line(xmin, v, xmax, v, size, color)
		c.Text(xmin, v, incr*.3, strconv.FormatFloat(v, 'g', -1, 64), color, op)
	}
}
