package gpdf

import (
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
	if c.CustomFont == nil {
		c.Page.AddTextColor(s, x, y, c.StdFont, size, creator.Color{R: color.R, G: color.G, B: color.B})
	} else {
		c.Page.AddTextCustomFontColor(s, x, y, c.CustomFont, size, creator.Color{R: color.R, G: color.G, B: color.B})
	}
}

func (c *Canvas) AbsBText(x, y, size float64, s string, fillcolor string) {
	c.AbsText(x, y, size, s, fillcolor)
}

func (c *Canvas) AbsCText(x, y, size float64, s string, fillcolor string) {
	if c.CustomFont == nil {
		c.AbsText(x, y, size, s, fillcolor)
		return
	}
	w := c.CustomFont.MeasureString(s, size)
	c.AbsText(x-(w/2), y, size, s, fillcolor)
}

func (c *Canvas) AbsEText(x, y, size float64, s string, fillcolor string) {
	if c.CustomFont == nil {
		c.AbsText(x, y, size, s, fillcolor)
		return
	}
	w := c.CustomFont.MeasureString(s, size)
	c.AbsText(x-w, y, size, s, fillcolor)
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
