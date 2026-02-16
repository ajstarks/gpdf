package gpdf

import (
	"math"

	"github.com/coregx/gxpdf/creator"
)

func (c *Canvas) AbsLine(x0, y0, x1, y1, size float64, strokecolor string) {
	clr, o := colorop(strokecolor)
	color := ColorLookup(clr)
	color.A = o / 100
	c.Page.DrawLine(x0, y0, x1, y1,
		&creator.LineOptions{
			Color: creator.Color{
				R: color.R,
				G: color.G,
				B: color.B,
			},
			Opacity: &color.A,
			Width:   size},
	)
}

func (c *Canvas) AbsCircle(x, y, r float64, fillcolor string) {
	color := ColorLookup(fillcolor)
	c.Page.DrawCircle(x, y, r, &creator.CircleOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A},
	)
}

func (c *Canvas) AbsEllipse(x, y, w, h float64, fillcolor string) {
	color := ColorLookup(fillcolor)
	c.Page.DrawEllipse(x, y, w, h, &creator.EllipseOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A})
}

func (c *Canvas) AbsCornerRect(x, y, w, h float64, fillcolor string) {
	color := ColorLookup(fillcolor)
	c.Page.DrawRect(x, y, w, h, &creator.RectOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A},
	)
}

func (c *Canvas) AbsCenterRect(x, y, w, h float64, fillcolor string) {
	c.AbsCornerRect(x-w/2, y-h/2, w, h, fillcolor)
}

func (c *Canvas) AbsText(x, y, size float64, s string, fillcolor string) {
	color := ColorLookup(fillcolor)
	c.Page.AddTextColor(s, x, y, c.Font, size, creator.Color{R: color.R, G: color.G, B: color.B})
}

func (c *Canvas) AbsPolygon(x, y []float64, fillcolor string) {
	lx := len(x)
	if lx != len(y) {
		return
	}
	coords := make([]creator.Point, lx)
	for i := range lx {
		coords[i].X = x[i]
		coords[i].Y = y[i]
	}
	color := ColorLookup(fillcolor)
	c.Page.DrawPolygon(coords, &creator.PolygonOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A},
	)
}

func (c *Canvas) AbsCubicBezier(bx, by, c0x, c0y, c1x, c1y, ex, ey, size float64, strokecolor string) {
	color := ColorLookup(strokecolor)
	bpoints := []creator.BezierSegment{
		{
			Start: creator.Point{X: bx, Y: by},
			C1:    creator.Point{X: c0x, Y: c0y},
			C2:    creator.Point{X: c1x, Y: c1y},
			End:   creator.Point{X: ex, Y: ey},
		},
	}
	c.Page.DrawBezierCurve(bpoints, &creator.BezierOptions{
		Color: creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		}, Width: size})
}

func deg2rad(d float64) float64 {
	return (360 - d) * (math.Pi / 180)
}

func (c *Canvas) AbsArc(x, y, rx, ry, angle1, angle2, size float64, strokecolor string) {
	/*
	   const n = 16

	   	for i := range n {
	   		p1 := float64(i+0) / n
	   		p2 := float64(i+1) / n
	   		a1 := angle1 + (angle2-angle1)*p1
	   		a2 := angle1 + (angle2-angle1)*p2
	   		x0 := x + rx*math.Cos(a1)
	   		y0 := y + ry*math.Sin(a1)
	   		x1 := x + rx*math.Cos((a1+a2)/2)
	   		y1 := y + ry*math.Sin((a1+a2)/2)
	   		x2 := x + rx*math.Cos(a2)
	   		y2 := y + ry*math.Sin(a2)
	   		cx := 2*x1 - x0/2 - x2/2
	   		cy := 2*y1 - y0/2 - y2/2

	   		c.AbsCubicBezier(x0, y0, cx, cy, cx, cy, x2, y2, size, strokecolor)

	   }
	*/
}

func (c *Canvas) AbsImage(x, y, w, h float64, name string) error {
	img, err := creator.LoadImage(name)
	if err != nil {
		return err
	}
	if h == 0 { // scaled image if height is zero
		imw := float64(img.Width())
		imh := float64(img.Height())
		scale := w / 100
		w = imw * scale
		h = imh * scale
	}
	err = c.Page.DrawImage(img, x, y, w, h)
	if err != nil {
		return err
	}
	return nil
}

func (c *Canvas) AbsCenterImage(x, y, w, h float64, name string) error {
	img, err := creator.LoadImage(name)
	if err != nil {
		return err
	}
	imw := float64(img.Width())
	imh := float64(img.Height())
	if h == 0 { // scaled image if height is zero
		scale := w / 100
		w = imw * scale
		h = imh * scale
	}
	x -= w / 2
	y -= h / 2
	err = c.Page.DrawImage(img, x, y, w, h)
	if err != nil {
		return err
	}
	return nil
}
