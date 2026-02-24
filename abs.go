package gpdf

import (
	"github.com/coregx/gxpdf/creator"
)

func (c *Canvas) AbsLine(x0, y0, x1, y1, size float64, color creator.ColorRGBA) {
	//println("AbsLine", color.A)
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

func (c *Canvas) AbsCircle(x, y, r float64, color creator.ColorRGBA) {
	// println("AbsCircle", color.A)
	c.Page.DrawCircle(x, y, r, &creator.CircleOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A},
	)
}

func (c *Canvas) AbsEllipse(x, y, w, h float64, color creator.ColorRGBA) {
	//println("AbsEllipse", x, y, w, h, color.R, color.G, color.B, color.A)
	c.Page.DrawEllipse(x, y, w, h, &creator.EllipseOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A})
}

func (c *Canvas) AbsCornerRect(x, y, w, h float64, color creator.ColorRGBA) {
	// println("AbsCornerRect", color.A)
	c.Page.DrawRect(x, y, w, h, &creator.RectOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A},
	)
}

func (c *Canvas) AbsGradRect(x, y, w, h float64, color1, color2 creator.ColorRGBA, pct float64) {
	var c1, c2 creator.Color
	c1.R, c1.G, c1.B = color1.R, color1.G, color1.B
	c2.R, c2.G, c2.B = color2.R, color2.G, color2.B
	grad := creator.NewLinearGradient(x, y, x+(w*pct), y+(h*pct))
	grad.Type = creator.GradientTypeLinear
	grad.AddColorStop(0, c1)
	grad.AddColorStop(1, c2)
	c.Page.DrawRect(x, y, w, h, &creator.RectOptions{FillGradient: grad})
}

func (c *Canvas) AbsCenterRect(x, y, w, h float64, color creator.ColorRGBA) {
	c.AbsCornerRect(x-w/2, y-h/2, w, h, color)
}

func (c *Canvas) AbsText(x, y, size float64, s string, color creator.ColorRGBA) {
	// println("AbsText", color.A)
	if c.CustomFont == nil {
		c.Page.AddTextColorAlpha(s, x, y, c.StdFont, size, creator.Color{R: color.R, G: color.G, B: color.B}, color.A)
	} else {
		c.Page.AddTextCustomFontColorAlpha(s, x, y, c.CustomFont, size, creator.Color{R: color.R, G: color.G, B: color.B}, color.A)
	}
}

func (c *Canvas) AbsBText(x, y, size float64, s string, fillcolor creator.ColorRGBA) {
	c.AbsText(x, y, size, s, fillcolor)
}

func (c *Canvas) AbsCText(x, y, size float64, s string, fillcolor creator.ColorRGBA) {
	if c.CustomFont == nil {
		c.AbsText(x, y, size, s, fillcolor)
		return
	}
	w := c.CustomFont.MeasureString(s, size)
	c.AbsText(x-(w/2), y, size, s, fillcolor)
}

func (c *Canvas) AbsEText(x, y, size float64, s string, fillcolor creator.ColorRGBA) {
	if c.CustomFont == nil {
		c.AbsText(x, y, size, s, fillcolor)
		return
	}
	w := c.CustomFont.MeasureString(s, size)
	c.AbsText(x-w, y, size, s, fillcolor)
}

func (c *Canvas) AbsRText(x, y, size, angle float64, s string, color creator.ColorRGBA) {
	if c.CustomFont == nil {
		c.Page.AddTextColorRotatedAlpha(s, x, y, c.StdFont, size, creator.Color{R: color.R, G: color.G, B: color.B}, angle, color.A)
	} else {
		c.Page.AddTextCustomFontColorRotatedAlpha(s, x, y, c.CustomFont, size, creator.Color{R: color.R, G: color.G, B: color.B}, angle, color.A)
	}
}

// whitespace determines if a rune is whitespace
func whitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}

func (c *Canvas) AbsPolygon(x, y []float64, color creator.ColorRGBA) {
	// println("AbsPolygon", color.A)
	lx := len(x)
	if lx < 3 || lx != len(y) {
		return
	}
	coords := make([]creator.Point, lx)
	for i := range lx {
		coords[i].X = x[i]
		coords[i].Y = y[i]
	}
	c.Page.DrawPolygon(coords, &creator.PolygonOptions{
		FillColor: &creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		},
		Opacity: &color.A},
	)
}

func (c *Canvas) AbsQuadBezier(bx, by, cx, cy, ex, ey, size float64, color creator.ColorRGBA) {
	// println("AbsQuad Bezier", color.A)
	bpoints := []creator.QuadBezierSegment{
		{
			Start:   creator.Point{X: bx, Y: by},
			Control: creator.Point{X: cx, Y: cy},
			End:     creator.Point{X: ex, Y: ey},
		},
	}
	c.Page.DrawQuadBezierCurve(bpoints, &creator.BezierOptions{
		Color: creator.Color{
			R: color.R,
			G: color.G,
			B: color.B,
		}, Width: size, Opacity: &color.A})
}

func (c *Canvas) AbsCubicBezier(bx, by, c0x, c0y, c1x, c1y, ex, ey, size float64, color creator.ColorRGBA) {
	// println("AbsCubic Bezier", color.A)
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
		}, Width: size, Opacity: &color.A})
}

func (c *Canvas) AbsImage(x, y, w, h float64, name string) error {
	//println(name, x, y, w, h)
	img, err := creator.LoadImage(name)
	if err != nil {
		return err
	}
	err = c.Page.DrawImage(img, x, y, w, h)
	if err != nil {
		return err
	}
	return nil
}

func (c *Canvas) AbsCenterImageName(x, y, w, h float64, name string) error {
	//println(name, x, y, w, h)
	img, err := creator.LoadImage(name)
	if err != nil {
		return err
	}
	x -= w / 2
	y -= h / 2
	err = c.Page.DrawImage(img, x, y, w, h)
	if err != nil {
		return err
	}
	return nil
}

func (c *Canvas) AbsCenterImage(x, y, w, h float64, img *creator.Image) error {
	x -= w / 2
	y -= h / 2
	err := c.Page.DrawImage(img, x, y, w, h)
	if err != nil {
		return err
	}
	return nil
}
