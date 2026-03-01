// gpdf generates PDF using a %-based coordinate system, with high-level functions for page elements
// convenience functions
package gpdf

import "math"

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
