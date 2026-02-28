package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/ajstarks/deck"
	"github.com/ajstarks/gpdf"
)

const (
	mm2pt        = 2.83464 // mm to pt conversion
	letterWidth  = 792
	letterHeight = 612
	linespacing  = 1.8
	listspacing  = 2.0
	listwrap     = 95.0
	defaultColor = "rgb(128,128,128)"
)

// PageDimen describes page dimensions
// the unit field is used to convert to pt.
type PageDimen struct {
	width, height, unit float64
}

// command line options
type options struct {
	sansfont   string
	serifont   string
	monofont   string
	symbolfont string
	layers     string
	pages      string
	pagesize   string
	output     string
	fontdir    string
	author     string
	title      string
	subject    string
	gridpct    float64
}

var (
	opts options

	// convert tabs to spaces
	codemap = strings.NewReplacer("\t", "    ")

	// fontmap maps generic font names to specific implementation names
	fontmap = map[string]gpdf.CF{}

	// cache for images
	imagecache = map[string]gpdf.Pimage{}

	// pagemap defines page dimensions
	pagemap = map[string]PageDimen{
		"Letter":     {792, 612, 1},
		"Legal":      {1008, 612, 1},
		"Tabloid":    {1224, 792, 1},
		"ArchA":      {864, 648, 1},
		"Widescreen": {1152, 648, 1},
		"4R":         {432, 288, 1},
		"Index":      {360, 216, 1},
		"A2":         {420, 594, mm2pt},
		"A3":         {420, 297, mm2pt},
		"A4":         {297, 210, mm2pt},
		"A5":         {210, 148, mm2pt},
	}
)

// includefile reads the content of a file into a string
func includefile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return ""
	}
	return codemap.Replace(string(data))
}

// textlines shows a series of lines of text
func textlines(canvas *gpdf.Canvas, x, y, size float64, s []string, c string) {
	yp := y
	for _, t := range s {
		canvas.Text(x, yp, size, t, c)
		yp -= linespacing * size
	}
}

// pagedim returns the page dimensions
func pagedim(s string) (float64, float64) {
	// try lookup first...
	v, ok := pagemap[s]
	if ok {
		return v.width, v.height
	}
	// lookup fails, try W<sep>H
	d := strings.FieldsFunc(s, func(c rune) bool { return !unicode.IsNumber(c) })
	if len(d) != 2 {
		return letterWidth, letterHeight
	}
	w, err := strconv.ParseFloat(d[0], 64)
	if err != nil {
		return letterWidth, letterHeight
	}
	h, err := strconv.ParseFloat(d[1], 64)
	if err != nil {
		return letterWidth, letterHeight
	}
	return w, h
}

// setfontdir determines the font directory:
// if the string argument is non-empty, use that, otherwise
// use the contents of the DECKFONT environment variable,
// if that is not set, or empty, use $HOME/deckfonts
func setfontdir(s string) string {
	if len(s) > 0 {
		return s
	}
	envdef := os.Getenv("DECKFONTS")
	if len(envdef) > 0 {
		return envdef
	}
	return path.Join(os.Getenv("HOME"), "deckfonts")
}

// pagerange returns the begin and end using a "-" string
func pagerange(s string) (int, int) {
	p := strings.Split(s, "-")
	if len(p) != 2 {
		return 0, 0
	}
	b, berr := strconv.Atoi(p[0])
	e, err := strconv.Atoi(p[1])
	if berr != nil || err != nil {
		return 0, 0
	}
	if b > e {
		return 0, 0
	}
	return b, e
}

// setopacity sets the alpha value:
// 0 == default value (opaque)
// -1 == fully transparent
// > 0 set opacity percent
func setopacity(v float64) float64 {
	var o float64
	switch {
	case v < 0:
		o = 0
	case v > 0:
		o = v
	case v == 0:
		o = 100
	}
	return o
}

func setmetadata(canvas *gpdf.Canvas) {
	if len(opts.author) > 0 {
		canvas.Creator.SetAuthor(opts.author)
	}
	if len(opts.title) > 0 {
		canvas.Creator.SetTitle(opts.title)
	}
	if len(opts.subject) > 0 {
		canvas.Creator.SetSubject(opts.subject)
	}
}

// cacheimages caches image objects for faster access
func cacheimages(c *gpdf.Canvas, d deck.Deck) {
	for i := range len(d.Slide) {
		slide := d.Slide[i]
		for j := range slide.Image {
			iname := slide.Image[j].Name
			img, err := c.LoadImage(iname)
			if err != nil {
				continue
			}
			_, ok := imagecache[iname]
			if !ok {
				imagecache[iname] = img
			}
		}
	}
}

// dimage processes deck images
func dimage(canvas *gpdf.Canvas, img gpdf.Pimage, i deck.Image) {
	sc := 100.0
	if i.Scale > 0 {
		sc = i.Scale
	}
	cw := canvas.Width
	x, y := i.Xp, i.Yp
	fw, fh := float64(i.Width), float64(i.Height)
	nw, nh := img.Width, img.Height
	// scale the image by the specified percentage
	if i.Scale > 0 {
		fw *= (i.Scale / 100)
		fh *= (i.Scale / 100)
	}
	// scale the image to fit the canvas width
	if i.Autoscale == "on" && fw > cw {
		fh *= (cw / fw)
		fw = cw
	}
	// scale the image to a percentage of the canvas width
	if i.Height == 0 && i.Width > 0 {
		if nh > 0 {
			imscale := (fw / 100) * cw
			fw = imscale
			fh = imscale / (nw / nh)
		}
	}
	canvas.Image(x, y, fw, fh, img)

	// process captions
	if len(i.Caption) > 0 {
		if i.Font == "" {
			i.Font = "sans"
		}
		if i.Sp == 0 {
			i.Sp = 1.8
		}
		c := i.Color
		canvas.CustomFont = fontmap[i.Font]
		ih := (float64(i.Height/2) / canvas.Height) * sc
		cx := i.Xp
		cy := (i.Yp - ih) - (i.Sp * 1.5)
		cs := i.Sp
		canvas.CText(cx, cy, cs, i.Caption, c)
	}
}

// line makes lines
func line(canvas *gpdf.Canvas, l deck.Line) {
	if l.Color == "" {
		l.Color = defaultColor
	}
	canvas.Line(l.Xp1, l.Yp1, l.Xp2, l.Yp2, l.Sp, l.Color, setopacity(l.Opacity))
}

// rect makes rectangles and squares
func rect(canvas *gpdf.Canvas, r deck.Rect) {
	if r.Color == "" {
		r.Color = defaultColor
	}
	c := r.Color
	op := setopacity(r.Opacity)
	x, y, w, h := r.Xp, r.Yp, r.Wp, r.Hp
	if r.Hr == 100 {
		canvas.Square(x, y, w, c, op)
	} else {
		canvas.Rect(x, y, w, h, c, op)
	}
}

// ellipse makes ellipses and circles
func ellipse(canvas *gpdf.Canvas, e deck.Ellipse) {
	if e.Color == "" {
		e.Color = defaultColor
	}
	c := e.Color
	op := setopacity(e.Opacity)
	x, y, w, h := e.Xp, e.Yp, e.Wp, e.Hp
	// println("ellipse", x, y, w, h)
	if e.Hr == 100 {
		canvas.Circle(x, y, w/2, c, op)
	} else {
		canvas.Ellipse(x, y, w/2, h/2, c, op)
	}
}

// arc makes arcs
func arc(canvas *gpdf.Canvas, a deck.Arc) {
	if a.Color == "" {
		a.Color = defaultColor
	}
	c := a.Color
	op := setopacity(a.Opacity)
	canvas.Arc(a.Xp, a.Yp, a.Wp, a.Hp, a.A1, a.A2, a.Sp, c, op)
}

// curve makea a quad bezier curve
func curve(canvas *gpdf.Canvas, curve deck.Curve) {
	if curve.Color == "" {
		curve.Color = defaultColor
	}
	c := curve.Color
	op := setopacity(curve.Opacity)
	x1, y1 := curve.Xp1, curve.Yp1
	x2, y2 := curve.Xp2, curve.Yp2
	x3, y3 := curve.Xp3, curve.Yp3
	sw := curve.Sp
	canvas.QuadCurve(x1, y1, x2, y2, x3, y3, sw, c, op)
}

// poly makes a filled polygon
func poly(canvas *gpdf.Canvas, p deck.Polygon) {
	xs := strings.Split(p.XC, " ")
	ys := strings.Split(p.YC, " ")
	if len(xs) != len(ys) {
		return
	}
	if len(xs) < 3 || len(ys) < 3 {
		return
	}
	xp := make([]float64, len(xs))
	yp := make([]float64, len(ys))
	for i := range xs {
		x, err := strconv.ParseFloat(xs[i], 64)
		if err != nil {
			xp[i] = 0
		} else {
			xp[i] = x
		}
		y, err := strconv.ParseFloat(ys[i], 64)
		if err != nil {
			yp[i] = 0
		} else {
			yp[i] = y
		}
	}
	if p.Color == "" {
		p.Color = defaultColor
	}
	canvas.Polygon(xp, yp, p.Color, setopacity(p.Opacity))
}

// dtext processes text
func dtext(canvas *gpdf.Canvas, t deck.Text) {
	if t.Font == "" {
		t.Font = "sans"
	}
	x, y, w, ts := t.Xp, t.Yp, t.Wp, t.Sp
	c := t.Color
	op := 100.0
	if t.Opacity > 0 {
		op = t.Opacity
	}
	canvas.CustomFont = fontmap[t.Font]

	s := t.Tdata
	if t.Type == "block" {
		canvas.TextWrapStrict(x, y, w, ts, 1.2, s, c, op)
		return
	}
	if len(t.File) > 0 {
		tl := strings.Split(includefile(t.File), "\n")
		if t.Type == "code" {
			canvas.CustomFont = fontmap["mono"]
			ch := float64(len(tl)) * linespacing * float64(ts)
			canvas.CornerRect(x-ts, y+(ts*2), (t.Wp), (ch), "rgb(240,240,240)")
		}
		textlines(canvas, x, y, ts, tl, c)
		return
	}
	if t.Rotation > 0 {
		canvas.RText(x, y, ts, t.Rotation, s, c, op)
		return
	}
	switch t.Align {
	case "c", "middle", "mid", "center":
		canvas.CText(x, y, ts, s, c, op)
	case "e", "right", "end":
		canvas.EText(x, y, ts, s, c, op)
	default:
		canvas.Text(x, y, ts, s, c, op)
	}
}

// bullet draws a bullet for a list item.
func bullet(canvas *gpdf.Canvas, x, y, size float64, c string) {
	canvas.Circle(x-size, y+size/2, size/4, c)
}

// number adds a number for a list item.
func number(canvas *gpdf.Canvas, n int, x, y, size float64, c string) {
	canvas.EText(x-size/2, y, size, fmt.Sprintf("%d.", n+1), c)
}

// list processes lists
func list(canvas *gpdf.Canvas, list deck.List) {
	rotation := list.Rotation
	font := list.Font
	color := list.Color
	align := list.Align
	ltype := list.Type
	ts := list.Sp
	ls := list.Lp
	xp := list.Xp
	yp := list.Yp
	op := setopacity(list.Opacity)
	if font == "" {
		font = "sans"
	}
	if ls == 0 {
		ls = listspacing
	}
	defont := font
	canvas.CustomFont = fontmap[defont]
	for i, item := range list.Li {
		t := item.ListText
		if len(item.Color) > 0 {
			color = item.Color
		} else {
			color = list.Color
		}
		if len(item.Font) > 0 {
			canvas.CustomFont = fontmap[item.Font]
		} else {
			canvas.CustomFont = fontmap[defont]
		}
		if item.Opacity > 0 {
			op = setopacity(item.Opacity)
		}
		if ltype == "number" {
			number(canvas, i, xp, yp, ts, color)
		}
		if ltype == "bullet" {
			bullet(canvas, xp, yp, ts, color)
		}
		if align == "center" {
			canvas.CText(xp, yp, ts, t, color, op)
		} else {
			if rotation == 0 {
				canvas.Text(xp, yp, ts, t, color, op)
			} else {
				canvas.RText(xp, yp, ts, rotation, t, color, op)
			}
		}
		yp -= (ls * ts)
	}
}

// process slides
func process(slideNumber int, d deck.Deck, canvas *gpdf.Canvas) {
	slide := d.Slide[slideNumber]
	var bg string
	if slide.Bg == "" {
		bg = "white"
	} else {
		bg = slide.Bg
	}
	if slide.Fg == "" {
		slide.Fg = "black"
	}

	if len(slide.Gradcolor1) > 0 && len(slide.Gradcolor2) > 0 {
		if slide.GradPercent <= 0 || slide.GradPercent > 100 {
			slide.GradPercent = 100
		}
		canvas.GradRect(0, 0, 100, 100, slide.Gradcolor1, slide.Gradcolor2, slide.GradPercent)
	} else {
		canvas.Background(bg)
	}

	// process each element according to the layer list
	layerlist := strings.Split(opts.layers, ":")
	for il := range layerlist {
		switch layerlist[il] { // for each element type loop through each set
		case "image":
			for _, i := range slide.Image {
				img, ok := imagecache[i.Name]
				if !ok {
					continue
				}
				dimage(canvas, img, i)
			}
		case "text":
			for _, t := range slide.Text {
				if t.Color == "" {
					t.Color = slide.Fg
				}
				dtext(canvas, t)
			}
		case "list":
			for _, li := range slide.List {
				list(canvas, li)
			}
		case "ellipse":
			for _, e := range slide.Ellipse {
				ellipse(canvas, e)
			}
		case "line":
			for _, l := range slide.Line {
				line(canvas, l)
			}
		case "rect":
			for _, r := range slide.Rect {

				rect(canvas, r)
			}
		case "poly":
			for _, p := range slide.Polygon {
				poly(canvas, p)
			}
		case "arc":
			for _, a := range slide.Arc {
				arc(canvas, a)
			}
		case "curve":
			for _, c := range slide.Curve {
				curve(canvas, c)
			}
		}
	}
	// add a grid, if specified
	if opts.gridpct > 0 {
		canvas.LGrid(0, 100, 0, 100, 0.1, opts.gridpct, slide.Fg, 20)
	}
}

// loadDeckFont reads TTF files from the font directory
func loadDeckFont(c *gpdf.Canvas, name, file string) {
	f, err := c.LoadFontFile(path.Join(opts.fontdir, file) + ".ttf")
	if err != nil {
		fontmap[name] = nil
	} else {
		fontmap[name] = f
	}
}

// dodeck writes PDF to w, processing deck markup read from r
// the page dimensions are (cw, ch), processing slides begin-end
func dodeck(w io.Writer, r io.ReadCloser, cw, ch float64, begin, end int) {
	canvas, err := gpdf.SetupCanvas(cw, ch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	setmetadata(canvas)
	// load default fonts
	loadDeckFont(canvas, "sans", opts.sansfont)
	loadDeckFont(canvas, "serif", opts.serifont)
	loadDeckFont(canvas, "mono", opts.monofont)
	loadDeckFont(canvas, "symbol", opts.symbolfont)
	d, err := deck.ReadDeck(r, 0, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	cacheimages(canvas, d)
	for i := 0; i < len(d.Slide); i++ {
		if i >= begin-1 && i <= end {
			if i > 0 {
				canvas.NewPage(cw, ch)
			}
			process(i, d, canvas)
		}
	}
	canvas.Creator.WriteTo(w)
}

// main: process commandline options, perform i/o, and page setup
func main() {
	// parse command line options
	flag.StringVar(&opts.author, "author", "", "document author")
	flag.StringVar(&opts.title, "title", "", "document title")
	flag.StringVar(&opts.subject, "subject", "", "document subject")
	flag.StringVar(&opts.sansfont, "sans", "NotoSans-Regular", "sans font")
	flag.StringVar(&opts.monofont, "mono", "Inconsolata-Medium", "mono font")
	flag.StringVar(&opts.serifont, "serif", "Charter-Regular", "sans font")
	flag.StringVar(&opts.symbolfont, "symbol", "ZapfDingbats", "sans font")
	flag.StringVar(&opts.layers, "layers", "image:rect:ellipse:curve:arc:line:poly:text:list", "Layer order")
	flag.StringVar(&opts.pagesize, "pagesize", "Letter", "pagesize: WxH, or one of: Letter, Legal, Tabloid, A3, A4, A5, ArchA, 4R, Index, Widescreen")
	flag.StringVar(&opts.pages, "pages", "1-1000000", "page range (first-last)")
	flag.StringVar(&opts.fontdir, "fontdir", setfontdir(""), "directory for fonts")
	flag.StringVar(&opts.output, "o", "", "output file")
	flag.Float64Var(&opts.gridpct, "grid", 0, "grid size (0 for no grid)")
	flag.Parse()

	// set up I/O (stdin, stdout by default)
	var w io.Writer = os.Stdout
	var r io.ReadCloser = os.Stdin
	var err error

	// input
	files := flag.Args()
	if len(files) > 0 {
		r, err = os.Open(files[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
	// output
	if len(opts.output) > 0 {
		w, err = os.Create(opts.output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
	// set up pages
	cw, ch := pagedim(opts.pagesize)
	begin, end := pagerange(opts.pages)
	dodeck(w, r, cw, ch, begin, end)
}
