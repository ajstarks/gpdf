package gpdf

import (
	"github.com/coregx/gxpdf/creator"
)

// built-in pagesizes and fonts
const (
	Letter  = creator.Letter
	Legal   = creator.Legal
	Tabloid = creator.Tabloid
	A3      = creator.A3
	A4      = creator.A4
	A5      = creator.A5
	B4      = creator.B4

	Mono  = creator.Courier
	Sans  = creator.Helvetica
	Serif = creator.TimesRoman
)

type CF *creator.CustomFont

type Pimage struct {
	Image         *creator.Image
	Width, Height float64
}

// Canvas object
type Canvas struct {
	Page          *creator.Page
	Creator       *creator.Creator
	StdFont       creator.FontName
	CustomFont    *creator.CustomFont
	Width, Height float64
}

// SetupCanvas initializes the canvas object with a standard page size
func SetupCanvasStdSize(pagesize creator.PageSize) (*Canvas, error) {
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

// SetupCanvas nitializes the canvas object, with specified dimensions
func SetupCanvas(width, height float64) (*Canvas, error) {
	canvas := new(Canvas)
	c := creator.New()
	page, err := c.NewPageWithDimensions(width, height)
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
func (c *Canvas) LoadFontFile(path string) (*creator.CustomFont, error) {
	f, err := creator.LoadFont(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// SetFont specifes the current font
func (c *Canvas) SetFont(f *creator.CustomFont) {
	c.CustomFont = f
}

// LoadImage loads a named image into the Pimage struct
func (c *Canvas) LoadImage(name string) (Pimage, error) {
	var p Pimage
	img, err := creator.LoadImage(name)
	if err != nil {
		return p, err
	}
	p.Image = img
	p.Width = float64(img.Width())
	p.Height = float64(img.Height())
	return p, nil
}

// NewPage starts a new Page
func (c *Canvas) NewPage(width, height float64) error {
	var err error
	c.Page, err = c.Creator.NewPageWithDimensions(width, height)
	return err
}
