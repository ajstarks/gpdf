// gpdf generates PDF using a %-based coordinate system, with high-level functions for page elements
// convenience functions
package gpdf

import (
	"math"
	"strconv"
)

// MapRange maps a value between low1 and high1, return the corresponding value between low2 and high2
func MapRange(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

// PolarDegrees returns the Cartesian coordinates (x, y) from polar coordinates
// with compensation for canvas aspect ratio
// center at (cx, cy), radius r, and angle theta (degrees)
func (c *Canvas) PolarDegrees(cx, cy, r, theta float64) (float64, float64) {
	dt := theta * (math.Pi / 180)
	aspect := c.Width / c.Height
	px := r * math.Cos(dt)
	py := (r * aspect) * math.Sin(dt)
	return cx + px, cy + py
}

// Polar returns the Cartesian coordinates (x, y) from polar coordinates
// with compensation for canvas aspect ratio
// center at (cx, cy), radius r, and angle theta (radians)
func (c *Canvas) Polar(cx, cy, r, theta float64) (float64, float64) {
	aspect := c.Width / c.Height
	px := r * math.Cos(theta)
	py := (r * aspect) * math.Sin(theta)
	return cx + px, cy + py
}

// Coord shows the specified coordinate, using percentage-based coordinates
// the (x, y) label is above the point, with a label below
func (c *Canvas) Coord(x, y, size float64, s string, fillcolor string) {
	c.Square(x, y, size/2, fillcolor)
	b := []byte("(")
	b = strconv.AppendFloat(b, x, 'g', -1, 32)
	b = append(b, ',')
	b = strconv.AppendFloat(b, y, 'g', -1, 32)
	b = append(b, ')')
	c.CText(x, y+size, size, string(b), fillcolor)
	if len(s) > 0 {
		c.CText(x, y-(size*1.33), size*0.66, s, fillcolor)
	}
}
