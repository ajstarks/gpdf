package gpdf

import (
	"github.com/coregx/gxpdf/creator"
)

// built-in pagesizes
const (
	Letter  = creator.Letter
	Legal   = creator.Legal
	Tabloid = creator.Tabloid
	A3      = creator.A3
	A4      = creator.A4
	A5      = creator.A5
	B4      = creator.B4
)
const (
	Mono  = creator.Courier
	Sans  = creator.Helvetica
	Serif = creator.TimesRoman
)

// Canvas object
type Canvas struct {
	Page          *creator.Page
	Creator       *creator.Creator
	StdFont       creator.FontName
	CustomFont    *creator.CustomFont
	Width, Height float64
}

// SetupCanvas initialized the canvas object
func SetupCanvas(pagesize creator.PageSize) (*Canvas, error) {
	canvas := new(Canvas)
	c := creator.New()
	page, err := c.NewPageWithSize(pagesize)
	if err != nil {
		return nil, err
	}
	canvas.Creator = c
	canvas.Page = page
	canvas.StdFont = Sans
	canvas.CustomFont = nil
	canvas.Width, canvas.Height = page.Width(), page.Height()
	return canvas, nil
}

// LoadFont loads a custom font
func (c *Canvas) LoadFont(path string) error {
	f, err := creator.LoadFont(path)
	if err != nil {
		return err
	}
	c.CustomFont = f
	return nil
}

// NewPage starts a new Page
func (c *Canvas) NewPage(pagesize creator.PageSize) error {
	var err error
	c.Page, err = c.Creator.NewPageWithSize(pagesize)
	return err
}
